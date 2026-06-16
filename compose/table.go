package compose

import (
	"fltk"
)

type _Table struct {
	Rect
	_Component
	wrap *fltk.TableRow
}

func (t *_Table) _Render() {
	t.applyModifier()
	if t.wrap == nil {
		t.wrap = fltk.NewTableRow(t.x, t.y, t.width, t.height)
	}
}

func Table(m ...Modifier) Component {
	t := &_Table{}
	t.modifiers = m
	t.self = t
	return t
}
