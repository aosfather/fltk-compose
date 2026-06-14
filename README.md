# fltk-compose
Compose is a GUI framework that adopts the Compose pattern, with its underlying implementation based on FLTK.

## Why fltk
We chose FLTK because it is lightweight and relatively mature. Our goal is to be more practical and production-ready, rather than merely showcasing technical flair.

## Why compose
We adopted the Compose pattern to make UI development more intuitive and straightforward, allowing developers to focus their energy on implementing core application functionality.

## Example
```go
package main

import "github.com/aosfather/fltk-compose/compose"

func main() {
	//one line
	compose.App(compose.Point(10, 10),compose.Size(400, 400),compose.Title("Hello compose world")).Run()
}

```

## Hello World
```go
package main

import (
	"github.com/aosfather/fltk-compose/attr"
	"github.com/aosfather/fltk-compose/compose"
)

func main() {
	//We use the attrpackage to separate the composepackage, distinguishing attributes from components.
	compose.App(attr.Point(10, 10), attr.Size(400, 400),
		attr.Title("Hello compose world")).Layout(compose.Label(attr.Point(20, 20),
		attr.Size(100, 25), attr.Title("Hello")),
		compose.Button(attr.Point(20, 50), attr.Size(100, 25), attr.Title("Click it")),
	).Run()
}


```

## Basic Component
### component list
* Label: Displays static text or images for user guidance.
* Button: Triggers an action when clicked by the user.
* RadioBox: Allows single selection from a group of mutually exclusive options.
* CheckBox: Enables toggling individual boolean options on or off.
* ComboBox: Combines a dropdown list with an editable text field.
* List: Presents a scrollable collection of selectable items.
* Messages: text list.
* Input: Collects user data (variants: text, password, integer, float).

### commond attrible
Here are simple descriptions for the methods under the attr package 
* Point: Represents the X and Y coordinates of a component.
* Size: Defines the width (W) and height (H) of a component.
* Title: Sets the title text of a window or component.



## Event
Events are handled through the Eventmethod, which binds standard component events and delegates them to your custom event handler functions. This approach keeps your code clean and focused, abstracting away the low-level details so you can concentrate on defining application behavior.

The event handling interface follows this signature:
```go
func(sender compose.Component, data *compose.EventData)
```
This provides a clean, unified contract: senderidentifies the component that triggered the event, while datacarries all relevant event information. It decouples your business logic from specific widget implementations, keeping handlers consistent and easy to maintain.

## Bind
The binding mechanism streamlines your workflow by letting you focus purely on application logic. Once a component is created and bound, data exchange happens automatically through the bind object—no need to wrestle with the underlying component APIs. It abstracts away the complexity, making state synchronization feel seamless and intuitive.
