package nimble

import (
	"code.google.com/p/nimble-cube/core"
	"code.google.com/p/nimble-cube/dump"
)

type Autotable struct {
	out   dump.TableWriter
	data  RChan3
	every int
}

func NewAutotable(fname string, data RChan3, every int) *Autotable {
	r := new(Autotable)
	tags := []string{data[0].Tag(), data[1].Tag(), data[2].Tag()}
	units := []string{"?", "?", "?"} // TODO
	r.out = dump.NewTableWriter(core.OpenFile(core.OD+fname), tags, units)
	r.data = data
	r.every = every
	return r
}

func (r *Autotable) Run() {
	core.Log("running auto table")
	N := core.Prod(r.data.Mesh().Size())

	for i := 0; ; i++ {
		output := r.data.ReadNext(N) // TODO
		if i%r.every == 0 {
			i = 0
			for c := range output {
				sum := 0.
				list := output[c].Host()
				for j := range list {
					sum += float64(list[j])
				}
				sum /= float64(N)
				r.out.Data[c] = float32(sum)
			}
		}
		r.data.ReadDone()
		r.out.WriteData()
		r.out.Flush()
	}
}