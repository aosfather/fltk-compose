package compose

import (
	"fltk"
	"fmt"
	"strconv"
	"strings"
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

const (
	_S_Color          = 3
	_S_BColor         = 4
	_S_Size           = 2
	_S_Aligment       = 1
	_bold       uint8 = 1
	_italic     uint8 = 2
	_underline  uint8 = 8
	_middleline uint8 = 4
)

var (
	_aligments = []string{"", "@r", "@c"}
	_sizes     = []string{"", "@s", "@m", "@l"}
)

type styleText uint8

const (
	st_Font styleText = 1
	st_Size styleText = 2
)

type Style struct {
	values []byte
	size   []byte
	text   styleText
}

func NewStyle() *Style {
	return &Style{values: make([]byte, 5), size: make([]byte, 2)}
}
func (s *Style) Color(i byte) *Style {
	s.values[_S_Color] = i
	return s
}

func (s *Style) BColor(i byte) *Style {
	s.values[_S_BColor] = i
	return s
}

func (s *Style) Large() *Style {
	s.values[_S_Size] = 3
	return s
}

func (s *Style) Medium() *Style {
	s.values[_S_Size] = 2
	return s
}

func (s *Style) Small() *Style {
	s.values[_S_Size] = 1
	return s
}
func (s *Style) Bold() *Style {
	s.values[0] = s.values[0] | _bold
	return s
}
func (s *Style) Italic() *Style {
	s.values[0] = s.values[0] | _italic
	return s
}

func (s *Style) Underline() *Style {
	s.values[0] = s.values[0] | _underline
	return s
}
func (s *Style) Middleline() *Style {
	s.values[0] = s.values[0] | _middleline

	return s
}
func (s *Style) Center() *Style {
	s.values[_S_Aligment] = 2
	return s
}

func (s *Style) Right() *Style {
	s.values[_S_Aligment] = 1
	return s
}
func (s *Style) Font(i byte) *Style {
	s.text |= st_Font
	s.size[0] = i
	return s
}

func (s *Style) Size(i byte) *Style {
	s.text |= st_Size
	s.size[1] = i
	return s
}

func (s *Style) ToFormat() string {
	format := strings.Builder{}
	for index, v := range s.values {
		switch index {
		case _S_Aligment:
			format.WriteString(_aligments[v])
		case _S_Size:
			format.WriteString(_sizes[v])
		case _S_Color:
			if v == 0 {
				continue
			}
			format.WriteString("@C")
			format.Write([]byte(strconv.Itoa(int(v))))
		case _S_BColor:
			if v == 0 {
				continue
			}
			format.WriteString("@B")
			format.Write([]byte(strconv.Itoa(int(v))))
		case 0:
			if v&_bold != 0 {
				format.WriteString("@b")
			}
			if v&_italic != 0 {
				format.WriteString("@i")
			}
			if v&_middleline != 0 {
				format.WriteString("@-")
			}
			if v&_underline != 0 {
				format.WriteString("@u")
			}
		}
	}

	if s.text&st_Font != 0 {
		format.WriteString("@F")
		format.WriteString(strconv.Itoa(int(s.size[0])))
	}

	if s.text&st_Size != 0 {
		format.WriteString("@S")
		format.WriteString(strconv.Itoa(int(s.size[1])))
	}

	return format.String()
}
