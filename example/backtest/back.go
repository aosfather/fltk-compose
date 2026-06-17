package main

import (
	"github.com/aosfather/fltk-compose/attr"
	"github.com/aosfather/fltk-compose/compose"
)

func main() {
	b := Back{}
	b.AddModifier(attr.Point(10, 10), attr.Size(300, 400), attr.Title("QuantEye"))
	b.Init()
	b.Run()
}

type Back struct {
	compose.Application
	product  compose.BindObj[string]
	project  compose.BindObj[string]
	view     compose.BindObj[compose.Log]
	progress compose.BindObj[float64]
}

func (b *Back) Init() {
	tabs := compose.Tabs(attr.Point(0, 0), attr.Size(300, 400))

	tabs.NewTab("回测",
		compose.Label(attr.Point(10, 15), attr.Size(20, 10), attr.Title("产品")),
		compose.Input(compose.IM_Normal, attr.Point(40, 10), attr.Size(180, 25)).Bind(&b.product),
		compose.Label(attr.Point(10, 45), attr.Size(20, 10), attr.Title("工程")),
		compose.Input(compose.IM_Normal, attr.Point(40, 40), attr.Size(180, 25)).Bind(&b.project),
		compose.Button(attr.Point(200, 75), attr.Size(80, 20), attr.Title("开始回测")).Event(b.startBackTest),
		compose.LogView(attr.Point(10, 100), attr.Size(280, 260)).Bind(&b.view))
	tabs.NewTab("盯盘",
		compose.Progress(attr.Point(10, 20), attr.Size(280, 20), attr.Min(0), attr.Max(100)).Bind(&b.progress),
		compose.Slider(attr.Point(10, 60), attr.Size(200, 20), attr.Max(16.6)),
		compose.Spinner(attr.Point(10, 90), attr.Size(100, 25)))
	b.Layout(tabs)
}

func (b *Back) startBackTest(sender compose.Component, data *compose.EventData) {
	log := b.view.Get()
	log.Error("buy:%s\n", b.product.Get())
	log.Warn("sell:%s", b.product.Get())
	v := b.progress.Get()
	b.progress.Set(v + 5)
}
