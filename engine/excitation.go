package engine

import (
	"github.com/mumax/3/cuda"
	"github.com/mumax/3/data"
	"github.com/mumax/3/script"
	"reflect"
)

// An excitation, typically field or current,
// can be defined region-wise plus extra mask*multiplier terms.
type excitation struct {
	perRegion  VectorParam // Region-based excitation
	extraTerms []mulmask   // add extra mask*multiplier terms
}

type mulmask struct {
	mul  func() float64
	mask *data.Slice
}

func (e *excitation) init(name, unit, desc string) {
	e.perRegion.init(name+"_perRegion", unit, "(internal)")
	DeclLValue(name, e, desc)
}

func (e *excitation) addTo(dst *data.Slice) {
	if !e.perRegion.isZero() {
		cuda.RegionAddV(dst, e.perRegion.LUT(), regions.Gpu())
	}
	for _, t := range e.extraTerms {
		var mul float32 = 1
		if t.mul != nil {
			mul = float32(t.mul())
		}
		cuda.Madd2(dst, dst, t.mask, 1, mul)
	}
}

func (e *excitation) isZero() bool {
	return e.perRegion.isZero() && len(e.extraTerms) == 0
}

func (e *excitation) Get() (*data.Slice, bool) {
	buf := cuda.Buffer(e.NComp(), e.Mesh())
	cuda.Zero(buf)
	e.addTo(buf)
	return buf, true
}

// Add an extra maks*multiplier term to the excitation.
func (e *excitation) Add(mask *data.Slice, mul func() float64) {
	e.extraTerms = append(e.extraTerms, mulmask{mul, assureGPU(mask)})
}

func assureGPU(s *data.Slice) *data.Slice {
	if s.GPUAccess() {
		return s
	} else {
		return cuda.GPUCopy(s)
	}
}

// user script: has to be 3-vector
func (e *excitation) SetRegion(region int, value [3]float64) {
	e.perRegion.setRegion(region, value[:])
}

// for gui (nComp agnostic)
func (e *excitation) setRegion(region int, value []float64) {
	e.perRegion.setRegion(region, value)
}

// does not use extramask!
func (e *excitation) getRegion(region int) []float64 {
	return e.perRegion.getRegion(region)
}

func (e *excitation) TableData() []float64 {
	return e.perRegion.getRegion(0)
}

func (p *excitation) Region(r int) TableData {
	return p.perRegion.Region(r)
}

// needed for script

func (e *excitation) SetValue(v interface{}) {
	e.perRegion.SetValue(v) // allows function of time
}

func (e *excitation) Name() string            { return e.perRegion.Name() }
func (e *excitation) Unit() string            { return e.perRegion.Unit() }
func (e *excitation) NComp() int              { return e.perRegion.NComp() }
func (e *excitation) Mesh() *data.Mesh        { return &globalmesh }
func (e *excitation) Eval() interface{}       { return e }
func (e *excitation) Type() reflect.Type      { return reflect.TypeOf(new(excitation)) }
func (e *excitation) InputType() reflect.Type { return script.VectorFunction_t }
