package engine

import (
	"github.com/mumax/3/cuda"
	"log"
)

func init() {
	DeclFunc("Run", Run, "Run the simulation for a time in seconds")
	DeclFunc("Steps", Steps, "Run the simulation for a number of time steps")
	DeclFunc("PostStep", PostStep, "Set up a function to be executed after every time step")
	DeclFunc("RunWhile", RunWhile, "Run while condition function is true")
	DeclVar("t", &Time, "Total simulated time (s)")
	DeclROnly("Dt", &Solver.Dt_si, "Last solver time step (s)")
	DeclVar("MinDt", &Solver.MinDt, "Minimum time step the solver can take (s)")
	DeclVar("MaxDt", &Solver.MaxDt, "Maximum time step the solver can take (s)")
	DeclVar("MaxErr", &Solver.MaxErr, "Maximum error per step the solver can tolerate")
	DeclVar("Headroom", &Solver.Headroom, "Solver headroom")
	DeclVar("FixDt", &Solver.FixDt, "Enable/disable fixed time step (default: false)")
}

var (
	Solver   cuda.Heun
	Time     float64             // time in seconds  // todo: hide? setting breaks autosaves
	pause    bool                // set pause at any time to stop running after the current step
	postStep []func()            // called on after every time step
	Inject   = make(chan func()) // injects code in between time steps. Used by web interface.
)

// Run the simulation for a number of seconds.
func Run(seconds float64) {
	log.Println("run for", seconds, "s")
	stop := Time + seconds
	RunWhile(func() bool { return Time < stop })
}

// Run the simulation for a number of steps.
func Steps(n int) {
	log.Println("run for", n, "steps")
	stop := Solver.NSteps + n
	RunWhile(func() bool { return Solver.NSteps < stop })
}

// Runs as long as condition returns true.
func RunWhile(condition func() bool) {
	checkM() // TODO: move to failed solver step

	pause = false
	for condition() && !pause {
		select {
		default:
			step()
		// accept tasks form Inject channel
		case f := <-Inject:
			f()
		}
	}
	pause = true
}

func step() {
	Solver.Step()
	for _, f := range postStep {
		f()
	}
	DoOutput()
}

// Register function f to be called after every time step.
// Typically used, e.g., to manipulate the magnetization.
func PostStep(f func()) {
	postStep = append(postStep, f)
}

// inject code into engine and wait for it to complete.
func InjectAndWait(task func()) {
	ready := make(chan int)
	Inject <- func() { task(); ready <- 1 }
	<-ready
}
