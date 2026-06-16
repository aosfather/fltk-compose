package compose

import "fmt"

type Modifier func(target any)
type _Point struct {
	x, y int
}

func (p *_Point) setX(target any) {
	target.(Basic).X(p.x)
}

func (p *_Point) setY(target any) {
	target.(Basic).Y(p.y)
}

func (p *_Point) setPoint(target any) {
	c := target.(Basic)
	c.X(p.x)
	c.Y(p.y)
}

func (m *_Point) String() string {
	return fmt.Sprintf("(x:%v,y:%v)", m.x, m.y)
}

// size
func X(x int) Modifier {
	p := &_Point{x: x}
	return p.setX
}
func Y(y int) Modifier {
	p := &_Point{y: y}
	return p.setY
}
func Point(x, y int) Modifier {
	p := &_Point{x: x, y: y}
	return p.setPoint
}

type _Size struct {
	w, h int
}

func (s *_Size) setWidth(target any) {
	target.(Basic).Width(s.w)
}

func (s *_Size) setHeight(target any) {
	target.(Basic).Height(s.h)
}

func (s *_Size) setSize(target any) {
	c := target.(Basic)
	c.Width(s.w)
	c.Height(s.h)
}

func (s *_Size) String() string {
	return fmt.Sprintf("[w:%v,h:%v)", s.w, s.h)
}
func Height(h int) Modifier {
	s := &_Size{h: h}
	return s.setHeight
}

func Width(w int) Modifier {
	s := &_Size{w: w}
	return s.setWidth
}

func MaxWidth() Modifier {
	s := &_Size{w: -1}
	return s.setWidth
}

func MaxHeight() Modifier {
	s := &_Size{h: -1}
	return s.setHeight
}

func MaxSize() Modifier {
	s := &_Size{w: -1, h: -1}
	return s.setSize
}

func Size(w, h int) Modifier {
	s := &_Size{w: w, h: h}
	return s.setSize
}

type _titleAble interface {
	SetTitle(string)
}

func Title(t string) Modifier {
	return func(target any) {
		target.(_titleAble).SetTitle(t)
	}
}

func Options(op []string) Modifier {
	return func(target any) {
		target.(_Options).SetOptions(op)
	}
}

type _Widths interface {
	SetColumnWidths(...int)
}

func Widths(i ...int) Modifier {
	return func(target any) {
		target.(_Widths).SetColumnWidths(i...)
	}
}

type _TextStyleAble interface {
	setTextStyle(int, string)
}

func TextStyles(s ...*Style) Modifier {
	return func(target any) {
		if s == nil {
			return
		}
		t := target.(_TextStyleAble)
		for index, st := range s {
			t.setTextStyle(index, st.ToFormat())
		}
	}
}
