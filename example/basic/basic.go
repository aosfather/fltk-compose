package main

import (
	"github.com/aosfather/fltk-compose/ui"

	"github.com/aosfather/fltk-compose/compose"
)

func main() {
	button := compose.Button(compose.Point(10, 10),
		compose.Size(60, 25),
		compose.Title("点一下")).Event(func(s compose.Component, e *compose.EventData) {
		ui.ShowMessage("消息", "hahah")
	})
	app := compose.App(compose.Point(10, 10),
		compose.Size(400, 400),
		compose.Title("测试一下"))
	app.Children(button,
		compose.Label(compose.Point(0, 50),
			compose.Size(100, 25),
			compose.Title("测试")),
		compose.CheckBox(compose.Point(0, 80), compose.Size(100, 25), compose.Title("选择")),
		compose.RadioBox(compose.Point(0, 110), compose.Size(100, 25), compose.Title("选择2")),
		compose.Column(compose.Point(0, 130), compose.Size(100, 25)).Children(
			compose.Input(compose.IM_Pass, compose.Size(100, 25)),
			compose.Input(compose.IM_Normal, compose.Size(100, 25)),
			compose.Row().Children(compose.Input(compose.IM_Float, compose.Size(100, 25)), compose.Input(compose.IM_Int, compose.Size(100, 25))),
			compose.ComboBox(compose.Point(0, 270), compose.Size(100, 25), compose.Options([]string{"A", "B", "C"}))),
	)
	app.Run()
}
