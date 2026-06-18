package attr

import "github.com/aosfather/fltk-compose/compose"

type ModifierHold struct {
	modifiers []compose.Modifier
}

func M() *ModifierHold {
	return &ModifierHold{}
}
func Modifier() *ModifierHold {
	return &ModifierHold{}
}
func (mh *ModifierHold) add(m compose.Modifier) {
	mh.modifiers = append(mh.modifiers, m)
}
func (mh *ModifierHold) Modifiers() []compose.Modifier {
	return mh.modifiers
}

func (mh *ModifierHold) M() []compose.Modifier {
	return mh.modifiers
}

func (mh *ModifierHold) Point(x, y int) *ModifierHold {
	mh.add(compose.Point(x, y))
	return mh
}

func (mh *ModifierHold) Size(x, y int) *ModifierHold {
	mh.add(compose.Size(x, y))
	return mh
}

func (mh *ModifierHold) Title(t string) *ModifierHold {
	mh.add(compose.Title(t))
	return mh
}

func (mh *ModifierHold) Options(op ...string) *ModifierHold {
	mh.add(compose.Options(op))
	return mh
}
func (mh *ModifierHold) Min(m float64) *ModifierHold {
	mh.add(compose.Min(m))
	return mh
}

func (mh *ModifierHold) Max(m float64) *ModifierHold {
	mh.add(compose.Max(m))
	return mh
}
func (mh *ModifierHold) Widths(i ...int) *ModifierHold {
	mh.add(compose.Widths(i...))
	return mh
}
