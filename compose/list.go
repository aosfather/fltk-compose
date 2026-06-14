package compose

import "fltk"

type _Options interface {
	SetOptions([]string)
}

type _Option struct {
	Options []string
}

func (op *_Option) SetOptions(o []string) {
	op.Options = o
}

// option 选择框
type _ComboBox struct {
	Rect
	_Component
	_Option
	wrap     *fltk.Choice
	OnSelect EventHandle
}

func (b *_ComboBox) _Render() {
	b.applyModifier()
	if b.wrap == nil {
		b.wrap = fltk.NewChoice(b.x, b.y, b.width, b.height)
		for _, s := range b.Options {
			b.wrap.Add(s, b.selected)
		}
	}
}

func (b *_ComboBox) SetSelect(index int) {
	if b.wrap != nil {
		b.wrap.SetValue(index)
	}
}

func (b *_ComboBox) GetSelect() int {
	if b.wrap != nil {
		return b.wrap.Value()
	}
	return -1
}

func (b *_ComboBox) SetSelectText(t string) {
	if b.wrap != nil {
		for index, v := range b.Options {
			if v == t {
				b.wrap.SetValue(index + 1)
			}
		}

	}
}
func (b *_ComboBox) GetSelectText() string {
	index := b.GetSelect()
	if index >= 0 {
		return b.Options[index]
	}
	return ""
}

func (b *_ComboBox) selected() {
	if b.OnSelect != nil {
		data := NewEventData("select")
		data.Setalue("text", b.wrap.SelectedText())
		data.Setalue("value", b.wrap.Value())
		b.OnSelect(b, data)
	}
}

func (b *_ComboBox) Bind(o *BindObj[int]) *_ComboBox {
	if o != nil {
		o.getter = b.GetSelect
		o.setter = b.SetSelect
	}

	return b
}

func (b *_ComboBox) BindText(o *BindObj[string]) *_ComboBox {
	if o != nil {
		o.getter = b.GetSelectText
		o.setter = b.SetSelectText
	}

	return b
}

func (b *_ComboBox) Event(h EventHandle) *_ComboBox {
	b.OnSelect = h
	return b
}

func ComboBox(m ...Modifier) *_ComboBox {
	b := &_ComboBox{}
	b.self = b
	b.modifiers = m
	return b
}
