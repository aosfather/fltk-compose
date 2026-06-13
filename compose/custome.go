package compose

type CustomeComponent interface {
	Children(c ...Component) Component
	Render()
}
type _WrapComponent struct {
	wrap CustomeComponent
}

func (w *_WrapComponent) AddModifier(modifiers ...Modifier) {

}

func (w *_WrapComponent) Children(c ...Component) Component {
	w.wrap.Children(c...)
	return w
}
func (w *_WrapComponent) _Render() {
	w.wrap.Render()
}
func Wrap(c CustomeComponent) Component {
	return &_WrapComponent{wrap: c}
}
