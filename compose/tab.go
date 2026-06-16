package compose

import (
	"fltk"
)

type _TitlePanel struct {
	Rect
	_Title
	_Component
	wrap *fltk.Tile
}

func (t *_TitlePanel) _Render() {
	t.applyModifier()
	if t.wrap == nil {
		t.wrap = fltk.NewTile(t.x, t.y+15, t.width, t.height, t.title)
	}
	t.wrap.Begin()
	for _, c := range t.children {
		c.AddModifier(fromPoint(t.x, t.y+15))
		c._Render()
	}
	t.wrap.End()
}

func TitlePanel(m ...Modifier) *_TitlePanel {
	tp := &_TitlePanel{}
	tp.modifiers = m
	tp.self = tp
	return tp
}

type _Group struct {
	title    string
	children []Component
}
type _Tabs struct {
	Rect
	_Component
	wrap   *fltk.Tabs
	groups []_Group
}

func (t *_Tabs) NewTab(title string, c ...Component) *_Tabs {
	g := _Group{title: title}
	g.children = append(g.children, c...)
	t.groups = append(t.groups, g)
	return t
}

func (t *_Tabs) _Render() {
	t.applyModifier()
	if t.wrap == nil {
		t.wrap = fltk.NewTabs(t.x, t.y, t.width, t.height)
	}
	t.wrap.Begin()
	for _, g := range t.groups {
		group := fltk.NewGroup(t.x, t.y+20, t.width, t.height, g.title)
		group.Begin()
		for _, c := range g.children {
			c.AddModifier(fromPoint(t.x, t.y+20))
			c._Render()
		}
		group.End()
	}

	t.wrap.End()
}

func Tabs(m ...Modifier) *_Tabs {
	t := &_Tabs{}
	t.modifiers = m
	t.self = t
	return t
}
