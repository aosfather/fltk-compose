package compose

import "fltk"

type InputModel byte

const (
	IM_Normal InputModel = 0
	IM_Pass   InputModel = 4
	IM_Int    InputModel = 1
	IM_Float  InputModel = 2
)

// 输入框
type _Input struct {
	Rect
	_Component
	wrap *fltk.Input

	OnChange EventHandle
	model    InputModel
}

func (i *_Input) _Render() {
	i.applyModifier()
	if i.wrap == nil {
		switch i.model {
		case IM_Pass:
			i.wrap = &fltk.NewSecretInput(i.x, i.y, i.width, i.height).Input
		case IM_Int:
			i.wrap = &fltk.NewIntInput(i.x, i.y, i.width, i.height).Input
		case IM_Float:
			i.wrap = &fltk.NewFloatInput(i.x, i.y, i.width, i.height).Input
		default:
			i.wrap = fltk.NewInput(i.x, i.y, i.width, i.height)
		}

		i.wrap.SetCallback(func() {
			if i.OnChange != nil {
				data := NewEventData("input")
				data.Setalue("text", i.wrap.Value())
				i.OnChange(i, data)
			}
		})
	}
}

func (b *_Input) SetValue(s string) {
	if b.wrap != nil {
		b.wrap.SetValue(s)
		b.wrap.Redraw()
	}
}
func (b *_Input) Value() string {
	if b.wrap != nil {
		return b.wrap.Value()
	}
	return ""
}

func (b *_Input) Event(h EventHandle) *_Input {
	b.OnChange = h
	return b
}

func Input(input_type InputModel, m ...Modifier) *_Input {
	input := &_Input{model: input_type}
	input.self = input
	input.modifiers = m
	return input
}
