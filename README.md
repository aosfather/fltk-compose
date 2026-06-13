# fltk-compose
Compose is a GUI framework that adopts the Compose pattern, with its underlying implementation based on FLTK.

## Why fltk
We chose FLTK because it is lightweight and relatively mature. Our goal is to be more practical and production-ready, rather than merely showcasing technical flair.

## Why compose
We adopted the Compose pattern to make UI development more intuitive and straightforward, allowing developers to focus their energy on implementing core application functionality.

# Example
```go
package main

import "github.com/aosfather/fltk-compose/compose"

func main() {
	//one line
	compose.App(compose.Point(10, 10),compose.Size(400, 400),compose.Title("Hello compose world")).Run()
}

```

# Hello World
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

# Basic Component
## Label
## Button
## RadioBox
## CheckBox
## ComboBox
## List
## TextList
## Input (text、pass、int、float)

# Event
