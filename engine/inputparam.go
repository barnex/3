package engine

import (
	"github.com/mumax/3/data"
	"github.com/mumax/3/util"
)

// input parameter, settable by user
type inputParam struct {
	lut
	upd_reg   [NREGION]func() []float64 // time-dependent values
	timestamp float64                   // used not to double-evaluate f(t)
	children  []derived                 // derived parameters
	descr
}

type derived interface {
	invalidate()
}

func (p *inputParam) init(nComp int, name, unit string, children []derived) {
	p.lut.init(nComp, p)
	p.descr = descr{name, unit}
	p.children = children
}

func (p *inputParam) update() {
	if p.timestamp != Time {
		changed := false
		// update functions of time
		for r := 0; r < NREGION; r++ { // TODO: 0..maxreg
			updFunc := p.upd_reg[r]
			if updFunc != nil {
				p.bufset_(r, updFunc())
				changed = true
			}
		}
		p.timestamp = Time
		if changed {
			p.invalidate()
		}
	}
}

// set in one region
func (p *inputParam) setRegion(region int, v []float64) {
	p.setRegions(region, region+1, v)
}

// set in all regions
// TODO: check if we always start from 0
func (p *inputParam) setUniform(v []float64) {
	p.setRegions(0, NREGION, v)
}

// set in regions r1..r2(excl)
func (p *inputParam) setRegions(r1, r2 int, v []float64) {
	util.Argument(len(v) == len(p.cpu_buf))
	util.Argument(r1 < r2) // exclusive upper bound
	for r := r1; r < r2; r++ {
		p.upd_reg[r] = nil
		p.bufset_(r, v)
	}
	p.invalidate()
}

func (p *inputParam) bufset_(region int, v []float64) {
	for c := range p.cpu_buf {
		p.cpu_buf[c][region] = float32(v[c])
	}
}

func (p *inputParam) setFunc(r1, r2 int, f func() []float64) {
	util.Argument(r1 < r2) // exclusive upper bound
	for r := r1; r < r2; r++ {
		p.upd_reg[r] = f
	}
	p.invalidate()
}

// mark my GPU copy and my children as invalid (need update)
func (p *inputParam) invalidate() {
	p.gpu_ok = false
	for _, c := range p.children {
		c.invalidate()
	}
}

func (p *inputParam) getRegion(region int) []float64 {
	cpu := p.CpuLUT()
	v := make([]float64, p.NComp())
	for i := range v {
		v[i] = float64(cpu[i][region])
	}
	return v
}

func (p *inputParam) IsUniform() bool {
	cpu := p.CpuLUT()
	v1 := p.getRegion(0)
	for r := 1; r < regions.maxreg; r++ {
		for c := range v1 {
			if cpu[c][r] != float32(v1[c]) {
				return false
			}
		}
	}
	return true
}

func (p *inputParam) Mesh() *data.Mesh { return Mesh() }

// Table output

// Parameter TableData is region 0
func (p *inputParam) TableData() []float64 {
	return p.getRegion(0)
}

func (p *inputParam) Region(r int) TableData {
	return &selectRegion{p, r}
}

// TableData for specific region
type selectRegion struct {
	*inputParam
	region int
}

func (p *selectRegion) TableData() []float64 {
	return p.getRegion(p.region)
}
