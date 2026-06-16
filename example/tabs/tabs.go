package main

import (
	"github.com/aosfather/fltk-compose/attr"
	"github.com/aosfather/fltk-compose/compose"
)

func main() {
	app := compose.App(attr.Point(10, 10), attr.Size(300, 400), attr.Title("tabs"))
	app.Layout(compose.Tabs(attr.Point(15, 35), attr.Size(200, 200)).NewTab("tab1").NewTab("tab2",
		compose.Button(attr.Point(0, 0), attr.Size(100, 20), attr.Title("push2")))).Run()

}
