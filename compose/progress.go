package compose

import (
	"fltk"
	"strconv"
)

// spin
type _Spin struct {
	Rect
	_Component
	wrap *fltk.Spinner
}

func (s *_Spin) _Render() {
	s.applyModifier()
	if s.wrap == nil {
		s.wrap = fltk.NewSpinner(s.x, s.y, s.width, s.height)
	}
}

// progress
type _Progress struct {
	Rect
	_Component
	wrap *fltk.Progress
	max  float64
	min  float64
}

func (p *_Progress) _Render() {
	p.applyModifier()
	if p.wrap == nil {
		p.wrap = fltk.NewProgress(p.x, p.y, p.width, p.height)
		p.wrap.SetSelectionColor(fltk.BLUE)
		if p.min >= 0 {
			p.wrap.SetMinimum(p.min)
		}
		if p.max > 0 {
			p.wrap.SetMaximum(p.max)
		}

	}
}

func (p *_Progress) SetMax(v float64) {
	if v > 0 {
		p.max = v
		if p.min > p.max {
			p.min = p.max
		}
	}

}

func (p *_Progress) SetMin(v float64) {
	if v >= 0 {
		p.min = v
	}
}

func (p *_Progress) Bind(o *BindObj[float64]) *_Progress {
	if o != nil {
		o.getter = func() float64 { return p.wrap.Value() }
		o.setter = func(v float64) {
			if v > p.max {
				v = p.max
			}
			if v < p.min {
				v = p.min
			}
			p.wrap.SetValue(v)
			p.wrap.SetLabel(strconv.Itoa(int(p.wrap.Value())))
		}
	}
	return p
}

func Progress(m ...Modifier) *_Progress {
	p := &_Progress{}
	p.modifiers = m
	p.self = p
	return p
}
