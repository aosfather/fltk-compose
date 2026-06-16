package main

import (
	"github.com/aosfather/fltk-compose/attr"
	"github.com/aosfather/fltk-compose/compose"
)

func main() {
	compose.App(attr.Point(100, 100),
		attr.Size(300, 400)).Layout(
		compose.TitlePanel(attr.Point(0, 0), attr.Size(200, 300),
			attr.Title("test1")).Children(compose.Table(attr.Point(5, 0),
			attr.Size(200, 300)))).Run()
}
