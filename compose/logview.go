package compose

import (
	"fltk"
	"fmt"
)

type _LogView struct {
	Rect
	_Component
	wrap   *fltk.Browser
	styles []string
}

type _LogLevel byte

const (
	_DEBUG _LogLevel = 0
	_INFO  _LogLevel = 1
	_WARN  _LogLevel = 2
	_ERROR _LogLevel = 3
)

func (l *_LogView) init() {
	l.styles = make([]string, 4)
	l.styles[_DEBUG] = "@C1"
	l.styles[_INFO] = "@C10"
	l.styles[_WARN] = "@C2"
	l.styles[_ERROR] = "@C3"

}
func (l *_LogView) _Render() {
	l.init()
	l.applyModifier()
	if l.wrap == nil {
		l.wrap = fltk.NewBrowser(l.x, l.y, l.width, l.height)
	}
}

func (l *_LogView) setTextStyle(t int, style string) {
	if t < 0 || t > len(l.styles) {
		return
	}
	l.styles[t] = style
}

type Log interface {
	Debug(format string, text ...any)
	Info(format string, text ...any)
	Warn(format string, text ...any)
	Error(format string, text ...any)
	To(int)
	ToTop()
	ToEnd()
}

var _LogNullSetter = func(v Log) {}

func (l *_LogView) Bind(o *BindObj[Log]) *_LogView {
	if o != nil {
		o.getter = func() Log { return l }
		o.setter = _LogNullSetter //空函数
	}
	return l
}

func (l *_LogView) log(t _LogLevel, format string, text ...any) {
	v := fmt.Sprintf(format, text...)
	if l.wrap != nil {
		l.wrap.Add(l.styles[t] + v)
	}
}
func (l *_LogView) Debug(format string, text ...any) {
	l.log(_DEBUG, format, text...)
}

func (l *_LogView) Info(format string, text ...any) {
	l.log(_INFO, format, text...)
}

func (l *_LogView) Warn(format string, text ...any) {
	l.log(_WARN, format, text...)
}

func (l *_LogView) Error(format string, text ...any) {
	l.log(_ERROR, format, text...)
}

func (m *_LogView) ToEnd() {
	m.wrap.SetValue(m.wrap.Size())
}

func (m *_LogView) ToTop() {
	m.wrap.SetValue(1)
}

func (m *_LogView) To(index int) {
	if index < 0 {
		index = 0
	}
	m.wrap.SetValue(index)
}

func LogView(m ...Modifier) *_LogView {
	l := &_LogView{}
	l.modifiers = m
	l.self = l
	return l

}
