package compose

import (
	"fltk"
	"strconv"
)

type _valueRange struct {
	min   float64
	max   float64
	value float64
}

func (s *_valueRange) SetMax(v float64) {
	s.max = v
	if s.min > s.max {
		s.min = s.max
	}

}

func (s *_valueRange) SetMin(v float64) {
	s.min = v
	if s.min > s.max {
		s.max = s.min
	}
}

func (s *_valueRange) SetValue(v float64) {
	s.value = v
}

// spinner
type _Spinner struct {
	Rect
	_Component
	_valueRange
	wrap *fltk.Spinner
}

func (s *_Spinner) _Render() {
	s.applyModifier()
	if s.wrap == nil {
		s.wrap = fltk.NewSpinner(s.x, s.y, s.width, s.height)
		s.wrap.SetType(fltk.SPINNER_FLOAT_INPUT)
	}
}

func (s *_Spinner) Bind(o *BindObj[float64]) *_Spinner {
	if o != nil {
		o.getter = func() float64 { return s.wrap.Value() }
		o.setter = func(v float64) {

		}
	}
	return s
}

func Spinner(m ...Modifier) *_Spinner {
	s := &_Spinner{}
	s.modifiers = m
	s.self = s
	return s
}

// Slider
type _Slider struct {
	Rect
	_Component
	wrap *fltk.Slider
	_valueRange
}

func (s *_Slider) _Render() {
	s.applyModifier()
	if s.wrap == nil {
		s.wrap = fltk.NewSlider(s.x, s.y, s.width, s.height)
		s.wrap.SetType(fltk.HOR_NICE_SLIDER)
		s.wrap.SetValue(s.value)
		s.wrap.SetMinimum(s.min)
		s.wrap.SetMaximum(s.max)
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
	_valueRange
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
