package compose

import (
	"fltk"
)

type Rect struct {
	width  int
	height int
	x      int
	y      int
}

func (r *Rect) X(x int) {
	r.x = x
}

func (r *Rect) Y(y int) {
	r.y = y
}

func (r *Rect) Width(w int) {
	r.width = w
}

func (r *Rect) Height(h int) {
	r.height = h
}

func NewRect(x, y, w, h int) Rect {
	return Rect{x: x, y: y, width: w, height: h}
}

type _Title struct {
	title string
}

func (t *_Title) SetTitle(text string) {
	t.title = text
}

func (t *_Title) GetTitle() string {
	return t.title
}

type _Component struct {
	control   _Control
	parent    any
	self      any
	modifiers []Modifier
	children  []Component
}

func (c *_Component) AddModifier(modifiers ...Modifier) {
	c.modifiers = append(c.modifiers, modifiers...)
}

func (c *_Component) setSelf(self any) {
	c.self = self
}

func (c *_Component) applyModifier() {
	for _, m := range c.modifiers {
		m(c.self)
	}
}

func (c *_Component) addChild(child Component) {
	c.children = append(c.children, child)
}

func (c *_Component) Children(children ...Component) Component {
	for _, component := range children {
		c.addChild(component)
	}

	return c.self.(Component)
}

func (b *_Component) Disable() {
	if b.control != nil {
		b.control.Deactivate()
	}
}

func (b *_Component) Enable() {
	if b.control != nil {
		b.control.Activate()
	}
}

func (b *_Component) IsEnable() bool {
	if b.control != nil {
		b.control.IsActive()
	}
	return false
}

type _Window struct {
	Rect
	_Title
	_Component
	wrap *fltk.Window
}

func (w *_Window) _Render() {
	w.applyModifier()
	if w.wrap == nil {
		if w.width <= 0 {
			w.width = 800
		}
		if w.height <= 0 {
			w.height = 600
		}
		w.wrap = fltk.NewWindow(w.width, w.height, w.title)
	}
	w.wrap.Begin()
	for _, c := range w.children {
		if l, ok := c.(Layout); ok {
			l.SetParentRect(&Rect{x: 0, y: 0, width: w.width, height: w.height})
		}
		c._Render()
	}
	w.wrap.End()

	w.wrap.SetPosition(w.x, w.y)

}

func (w *_Window) Show() {
	if w.wrap != nil {
		w.wrap.Show()
	}
}

type Application struct {
	Rect
	_Title
	_Component
	mainWindow *_Window
}

func (a *Application) Run() {
	if a.mainWindow == nil {
		a.mainWindow = Window(a.modifiers...)
	}

	a.mainWindow._Render()
	a.mainWindow.Show()
	fltk.Run()
}

func (a *Application) Layout(c ...Component) *Application {
	if a.mainWindow == nil {
		a.mainWindow = Window(a.modifiers...)
	}
	for _, component := range c {
		a.mainWindow.addChild(component)
	}

	return a
}

func Window(m ...Modifier) *_Window {
	w := &_Window{}
	w.modifiers = m
	w.self = w
	return w
}

func App(m ...Modifier) *Application {
	app := &Application{}
	app.modifiers = m
	app.setSelf(app)
	return app
}
