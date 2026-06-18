package main

import (
	"github.com/aosfather/fltk-compose/attr"
	"github.com/aosfather/fltk-compose/compose"
)

func main() {
	b := Back{}
	b.AddModifier(attr.M().Point(10, 10).Size(300, 400).Title("QuantEye").M()...)
	b.Init()
	b.Run()
}

type Back struct {
	compose.Application
	product  compose.BindObj[string]
	project  compose.BindObj[string]
	view     compose.BindObj[compose.Log]
	progress compose.BindObj[float64]
	tv       compose.BindObj[compose.TableView]
}

func (b *Back) Init() {
	tabs := compose.Tabs(attr.M().Point(0, 0).Size(300, 400).M()...)

	tabs.NewTab("回测",
		compose.Label(attr.M().Point(10, 15).Size(20, 10).Title("产品").M()...),
		compose.Input(compose.IM_Normal, attr.M().Point(40, 10).Size(180, 25).M()...).Bind(&b.product),
		compose.Label(attr.M().Point(10, 45).Size(20, 10).Title("工程").M()...),
		compose.Input(compose.IM_Normal, attr.M().Point(40, 40).Size(180, 25).M()...).Bind(&b.project),
		compose.Button(attr.M().Point(200, 75).Size(80, 20).Title("开始回测").M()...).Event(b.startBackTest),
		compose.LogView(attr.M().Point(10, 100).Size(280, 260).M()...).Bind(&b.view))

	tabs.NewTab("盯盘",
		compose.Progress(attr.M().Point(10, 20).Size(280, 20).Min(0).Max(100).M()...).Bind(&b.progress),
		compose.Slider(attr.M().Point(10, 60).Size(200, 20).Max(16.6).M()...),
		compose.Spinner(attr.M().Point(10, 90).Size(100, 25).M()...),
		compose.Table(attr.M().Point(10, 120).Size(260, 200).Options("第一列", "第二列").M()...).Bind(&b.tv))
	b.Layout(tabs)
}

func (b *Back) startBackTest(sender compose.Component, data *compose.EventData) {
	log := b.view.Get()
	log.Error("buy:%s\n", b.product.Get())
	log.Warn("sell:%s", b.product.Get())
	v := b.progress.Get()
	b.progress.Set(v + 5)
	b.tv.Get().AddRow("测试", b.product.Get())
}
