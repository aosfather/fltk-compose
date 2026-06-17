package main

import (
	"github.com/aosfather/fltk-compose/attr"
	"github.com/aosfather/fltk-compose/compose"
)

func main() {
	b := Back{}
	b.AddModifier(attr.Point(10, 10), attr.Size(300, 400), attr.Title("回测"))
	b.Init()
	b.Run()
}

type Back struct {
	compose.Application
	product compose.BindObj[string]
	project compose.BindObj[string]
	view    compose.BindObj[compose.Log]
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
	b.Layout(tabs)
}

func (b *Back) startBackTest(sender compose.Component, data *compose.EventData) {
	log := b.view.Get()
	log.Error("buy:%s\n", b.product.Get())
	log.Warn("sell:%s", b.product.Get())
}
