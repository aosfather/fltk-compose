package compose

import (
	"fltk"
	"strconv"
)

// Slider
type _Slider struct {
	Rect
	_Component
	wrap         *fltk.Slider
	max          float64
	min          float64
	defaultValue float64
}

func (s *_Slider) _Render() {
	s.applyModifier()
	if s.wrap == nil {
		s.wrap = fltk.NewSlider(s.x, s.y, s.width, s.height)
		s.wrap.SetType(fltk.HOR_NICE_SLIDER)
		s.wrap.SetValue(s.defaultValue)
		s.wrap.SetMinimum(s.min)
		s.wrap.SetMaximum(s.max)
	}
}

func (s *_Slider) SetMax(v float64) {
	if v > 0 {
		s.max = v
		if s.min > s.max {
			s.min = s.max
		}
	}

}

func (s *_Slider) SetMin(v float64) {
	if v >= 0 {
		s.min = v
	}
}

func (s *_Slider) Bind(o *BindObj[float64]) *_Slider {
	if o != nil {
		o.getter = func() float64 { return s.wrap.Value() }
		o.setter = func(v float64) {

		}
	}
	return s
}

func Slider(m ...Modifier) *_Slider {
	s := &_Slider{}
	s.modifiers = m
	s.self = s
	return s
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
