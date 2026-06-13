package main

import (
	"github.com/aosfather/fltk-compose/attr"
	"github.com/aosfather/fltk-compose/compose"
)

func main() {
	compose.App(attr.Point(10, 10), attr.Size(400, 400),
		attr.Title("Hello compose world")).Layout(compose.Label(attr.Point(20, 20),
		attr.Size(100, 25), attr.Title("Hello")),
		compose.Button(attr.Point(20, 50), attr.Size(100, 25), attr.Title("Click it")),
	).Run()
}
