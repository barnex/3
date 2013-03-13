package engine

// File: Quant stores a physical quantity.
// Author: Arne Vansteenkiste

import (
	"code.google.com/p/mx3/data"
	"log"
)

// shared by Quant and Reader
type quant struct {
	buffer data.Slice           // stores the data
	lock   [data.MAX_COMP]mutex // protects buffer. mutex can be read or write type TODO: make slice, also for Slice
}

type Quant struct {
	quant
}

func QuantFromSlice(s *data.Slice) *Quant {
	N := s.Mesh().NCell()
	var lock [data.MAX_COMP]mutex
	nComp := s.NComp()
	for c := 0; c < nComp; c++ {
		lock[c] = newRWMutex(N)
	}
	return &Quant{quant{*s, lock}}
}

type Reader struct {
	quant
}

func (c *Quant) NewReader() *Reader {
	reader := (*c).quant // copy
	for i := range reader.lock {
		reader.lock[i] = (reader.lock[i]).(*rwMutex).MakeRMutex()
	}
	return &Reader{reader}
}

// NComp returns the number of components.
// 	scalar: 1
// 	vector: 3
// 	...
func (c *quant) NComp() int {
	return int(c.buffer.NComp())
}

func (c *quant) comp(i int) quant {
	return quant{*c.buffer.Comp(i), [data.MAX_COMP]mutex{c.lock[i]}}
}

func (c *Quant) Comp(i int) Quant {
	return Quant{c.comp(i)}
}

func (c *Reader) Comp(i int) Reader {
	return Reader{c.comp(i)}
}

func (c *quant) Mesh() *data.Mesh {
	return c.buffer.Mesh()
}

// Returns the data buffer without locking.
// To be used with extreme care.
func (c *quant) Data() *data.Slice {
	for i := 0; i < c.NComp(); i++ {
		if c.lock[i].isLocked() {
			log.Panic("quant unsafe data: is locked")
		}
	}
	return &c.buffer
}

// lock the next n elements of buffer.
func (c *quant) next() *data.Slice {
	//n := c.Mesh().NCell()
	c.lock[0].lockNext()
	a, b := c.lock[0].lockedRange()
	ncomp := c.NComp()
	for i := 1; i < ncomp; i++ {
		c.lock[i].lockNext()
		α, β := c.lock[i].lockedRange()
		if α != a || β != b {
			panic("chan: next: inconsistent state")
		}
	}
	return c.buffer.Slice(a, b)
}

// unlock the locked buffer range.
func (c *quant) done() {
	for i := 0; i < c.NComp(); i++ {
		c.lock[i].unlock()
	}
}

func (c *Quant) WriteNext() *data.Slice {
	return c.quant.next()
}

func (c *Quant) WriteDone() {
	c.quant.done()
}

func (c *Reader) ReadNext() *data.Slice {
	return c.quant.next()
}

func (c *Reader) ReadDone() {
	c.quant.done()
}