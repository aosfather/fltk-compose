package compose

import (
	"fltk"
)

// BUTTON
type _Button struct {
	Rect
	_Title
	_Component
	wrap        *fltk.Button
	eventHandle EventHandle
}

func (b *_Button) _Render() {
	b.applyModifier()
	if b.wrap == nil {
		b.wrap = fltk.NewButton(b.x, b.y, b.width, b.height, b.title)
		b.wrap.SetCallback(func() {
			if b.eventHandle != nil {
				b.eventHandle(b, nil)
			}
		})

		b.control = b.wrap
	}
}

func (b *_Button) Event(h EventHandle) *_Button {
	b.eventHandle = h
	return b
}

func Button(m ...Modifier) *_Button {
	b := &_Button{}
	b.self = b
	b.modifiers = m
	return b
}

//label

type _Label struct {
	Rect
	_Title
	_Component
	wrap        *fltk.Box
	eventHandle EventHandle
}

func (l *_Label) _Render() {
	l.applyModifier()
	if l.wrap == nil {
		l.wrap = fltk.NewBox(fltk.BORDER_FRAME, l.x, l.y, l.width, l.height, l.title)
		l.control = l.wrap
	}
}

func Label(m ...Modifier) *_Label {
	b := &_Label{}
	b.self = b
	b.modifiers = m
	return b
}

// checkbox
type _CheckBox struct {
	Rect
	_Title
	_Component
	wrap    *fltk.CheckButton
	OnCheck EventHandle
}

func (b *_CheckBox) _Render() {
	b.applyModifier()
	if b.wrap == nil {
		b.wrap = fltk.NewCheckButton(b.x, b.y, b.width, b.height, b.title)
		b.wrap.SetCallback(func() {
			if b.OnCheck != nil {
				data := NewEventData("check")
				data.Setalue("select", b.wrap.Value())
				b.OnCheck(b, data)
			}

		})
	}
}

func (b *_CheckBox) setValue(check bool) {
	b.wrap.SetValue(check)
}
func (b *_CheckBox) Value() bool {
	return b.wrap.Value()
}

func (b *_CheckBox) Bind(o *BindObj[bool]) *_CheckBox {
	if o != nil {
		o.setter = b.setValue
		o.getter = b.Value
	}
	return b
}

func CheckBox(m ...Modifier) *_CheckBox {
	b := &_CheckBox{}
	b.self = b
	b.modifiers = m
	return b
}

// readiobox
type _RadioBox struct {
	Rect
	_Title
	_Component
	wrap    *fltk.RadioRoundButton
	OnCheck EventHandle
}

func (b *_RadioBox) _Render() {
	b.applyModifier()
	if b.wrap == nil {
		b.wrap = fltk.NewRadioRoundButton(b.x, b.y, b.width, b.height, b.title)
		b.wrap.SetCallback(func() {
			if b.OnCheck != nil {
				data := NewEventData("check")
				data.Setalue("select", b.wrap.Value())
				b.OnCheck(b, data)
			}

		})
	}
}

func RadioBox(m ...Modifier) *_RadioBox {
	b := &_RadioBox{}
	b.self = b
	b.modifiers = m
	return b
}
