package compose

/*
*
* Composable
* fun MessageList(messages: List<String>) {
LazyColumn {
items(messages) { message ->
Text(message, modifier = Modifier.padding(8.dp))
}
}
}
*/
type ModifierAble interface {
	AddModifier(modifiers ...Modifier)
}

type Component interface {
	ModifierAble
	Children(c ...Component) Component
	_Render()
	// Event(int, EventHandle)
}

type Modifier func(target any)
type Composable func(modifiers ...Modifier) Component

/**
 * 能力
 */
type Basic interface {
	X(x int)
	Y(y int)
	Width(w int)
	Height(h int)
}

type _Control interface {
	Deactivate()
	IsActive() bool
	Activate()
}

type Control interface {
	Basic
	Disable()
	Enable()
	IsEnable() bool
}
