package main

import (
	"github.com/aosfather/fltk-compose/attr"
	"github.com/aosfather/fltk-compose/compose"
	"github.com/aosfather/fltk-compose/ui"
)

func main() {
	button := compose.Button(attr.Point(10, 10),
		attr.Size(60, 25),
		attr.Title("点一下")).Event(func(s compose.Component, e *compose.EventData) {
		ui.ShowMessage("消息", "hahah")
	})
	compose.App(attr.Point(10, 10),
		attr.Size(400, 400),
		attr.Title("测试一下")).Layout(button,
		compose.Label(attr.Point(0, 50),
			attr.Size(100, 25),
			attr.Title("测试")),
		compose.CheckBox(attr.Point(0, 80), attr.Size(100, 25), attr.Title("选择")),
		compose.RadioBox(attr.Point(0, 110), attr.Size(100, 25), attr.Title("选择2")),
		compose.Column(attr.Point(0, 130), attr.Size(100, 25)).Children(
			compose.Input(compose.IM_Pass, attr.Size(100, 25)),
			compose.Input(compose.IM_Normal, attr.Size(100, 25)),
			compose.Row().Children(compose.Input(compose.IM_Float, attr.Size(100, 25)), compose.Input(compose.IM_Int, compose.Size(100, 25))),
			compose.ComboBox(attr.Point(0, 270), attr.Size(100, 25), compose.Options([]string{"A", "B", "C"}))),
	).Run()
}
