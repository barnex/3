package gpu

import (
	"code.google.com/p/nimble-cube/core"
	"code.google.com/p/nimble-cube/dump"
	"code.google.com/p/nimble-cube/nimble"
	"github.com/barnex/cuda5/cu"
	"github.com/barnex/cuda5/safe"
	"math"
)

// Adaptive heun solver.
// TODO: now only for magnetization (because it normalizes)
// post-step hook?
type RK23 struct {
	dy0 [3]safe.Float32s // buffer dy/dt
	y   nimble.ChanN
	dy  nimble.RChanN
	solverCommon
	init   bool
	stream [3]cu.Stream
}

func NewRK23(y nimble.ChanN, dy_ nimble.ChanN, dt, multiplier float64) *RK23 {
	core.Assert(dt > 0 && multiplier > 0)
	dy := dy_.NewReader()
	dy0 := MakeVectors(y.BufLen()) // TODO: proper len?
	var w dump.TableWriter
	if core.DEBUG {
		w = dump.NewTableWriter(core.OpenFile(core.OD+"/debug_heun.table"),
			[]string{"t", "dt", "err"}, []string{"s", "s", y.Unit()})
	}
	return &RK23{dy0: dy0, y: y, dy: dy,
		solverCommon: solverCommon{dt_si: dt, dt_mul: multiplier, Maxerr: 1e-4, Headroom: 0.75, debug: w},
		stream:       stream3Create()}
}

// Run for a duration in seconds
func (e *RK23) Advance(seconds float64) {
	nimble.RunStack()
	core.Log("GPU heun solver:", seconds, "s")
	LockCudaThread()
	stop := e.time + seconds
	for e.time < stop {
		e.Step()
	}
	nimble.DashExit()

	if core.DEBUG {
		e.debug.Flush()
	}
}

// Run for a number of steps
func (e *RK23) Steps(steps int) {
	nimble.RunStack()
	core.Log("GPU heun solver:", steps, "steps")
	LockCudaThread()
	for s := 0; s < steps; s++ {
		e.Step()
	}
	nimble.DashExit()

	if core.DEBUG {
		e.debug.Flush()
	}
}

// Take one time step
func (e *RK23) Step() {
	n := e.y.Mesh().NCell()
	str := e.stream

	// Send out initial value
	if !e.init {
		// normalize initial magnetization
		M := Device3(e.y.UnsafeData())
		NormalizeSync(M, str[0])
		e.y.WriteNext(n)
		e.init = true
	}
	e.y.WriteDone()

	dy0 := e.dy0
	dt := float32(e.dt_si * e.dt_mul) // could check here if it is in float32 ranges
	core.Assert(dt > 0)

	// stage 1
	nimble.Clock.Send(e.time, true)
	dy := Device3(e.dy.ReadNext(n))
	y := Device3(e.y.WriteNext(n))
	{
		for i := 0; i < 3; i++ {
			Madd2Async(y[i], y[i], dy[i], 1, dt, str[i])
			dy0[i].CopyDtoDAsync(dy[i], str[i])
		}
		syncAll(str[:])
	}
	e.y.WriteDone()
	e.dy.ReadDone()

	// stage 2
	nimble.Clock.Send(e.time+e.dt_si, false)
	dy = Device3(e.dy.ReadNext(n))
	y = Device3(e.y.WriteNext(n))
	{
		err := MaxVecDiff(dy0[0], dy0[1], dy0[2], dy[0], dy[1], dy[2], str[0]) * float64(dt)
		if err == 0 {
			nimble.DashExit()
			core.Fatalf("heun: cannot adapt dt")
			// Note: err == 0 occurs when input is NaN (or time step massively too small).
		}

		if core.DEBUG {
			e.debug.Data[0], e.debug.Data[1], e.debug.Data[2] = float32(e.time), float32(e.dt_si), float32(err)
			e.debug.WriteData()
		}

		if err < e.Maxerr || e.dt_si <= e.Mindt { // mindt check to avoid infinite loop
			Madd3Async(y[0], y[0], dy[0], dy0[0], 1, 0.5*dt, -0.5*dt, str[0])
			Madd3Async(y[1], y[1], dy[1], dy0[1], 1, 0.5*dt, -0.5*dt, str[1])
			Madd3Async(y[2], y[2], dy[2], dy0[2], 1, 0.5*dt, -0.5*dt, str[2])
			syncAll(str[:])
			NormalizeSync(y, str[0])
			e.time += e.dt_si
			e.steps++
			e.adaptDt(math.Pow(e.Maxerr/err, 1./2.))
		} else {
			// undo.
			Madd2Async(y[0], y[0], dy0[0], 1, -dt, str[0])
			Madd2Async(y[1], y[1], dy0[1], 1, -dt, str[1])
			Madd2Async(y[2], y[2], dy0[2], 1, -dt, str[2])
			e.undone++
			e.adaptDt(math.Pow(e.Maxerr/err, 1./3.))
		}
		e.updateDash(err)
	}
	e.dy.ReadDone()
	// no writeDone() here.
}