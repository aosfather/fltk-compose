package compose

import (
	"fltk"
	"fmt"
)

// 消息列表
type _Messages struct {
	Rect
	_Component
	wrap    *fltk.Browser
	columns []int
}

func (m *_Messages) SetColumnWidths(w ...int) {
	m.columns = w
}
func (m *_Messages) AddString(layout string, text ...any) {
	v := fmt.Sprintf(layout, text...)
	if m.wrap != nil {
		m.wrap.Add(v)
	}

}
func (m *_Messages) _Render() {
	m.applyModifier()
	if m.wrap == nil {
		m.wrap = fltk.NewBrowser(m.x, m.y, m.width, m.height)
		if m.columns != nil {
			m.wrap.SetColumnWidths(m.columns...)
		}

	}
}

func (m *_Messages) ToEnd() {
	m.wrap.SetValue(m.wrap.Size())
}

func (m *_Messages) ToTop() {
	m.wrap.SetValue(1)
}

func (m *_Messages) To(index int) {
	if index < 0 {
		index = 0
	}
	m.wrap.SetValue(index)
}

// 空函数
var _NullSetter = func(v MessageList) {}

func (m *_Messages) Bind(o *BindObj[MessageList]) *_Messages {
	if o != nil {
		o.getter = func() MessageList { return m }
		o.setter = _NullSetter //空函数
	}
	return m
}

type MessageList interface {
	AddString(layout string, text ...any)
	To(int)
	ToTop()
	ToEnd()
}

func Messages(m ...Modifier) *_Messages {
	c := &_Messages{}
	c.self = c
	c.modifiers = m
	return c
}
