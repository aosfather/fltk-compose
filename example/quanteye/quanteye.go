package main

import (
	"fmt"
	"strings"

	"github.com/aosfather/fltk-compose/attr"
	"github.com/aosfather/fltk-compose/compose"
	"github.com/aosfather/fltk-compose/ui"
)

func main() {
	app := QuantEye{}
	app.Layout(attr.Size(400, 500), attr.Title("盯盘")).Run()
}

type QuantEye struct {
	application *compose.Application
	it_product  compose.BindObj[string]
	cc_val      compose.BindObj[int]
	it_project  compose.BindObj[string]
	l_result    compose.BindObj[string]
	cb_weixin   compose.BindObj[bool]
	weixin      *Weixin
	recentValue *Recent
}

func (q *QuantEye) Layout(m ...compose.Modifier) *compose.Application {
	q.application = compose.App(m...)
	q.application.Layout(compose.Label(attr.Point(10, 15), attr.Size(20, 10), attr.Title("产品")),
		compose.Input(compose.IM_Normal, attr.Point(40, 10), attr.Size(180, 25)).Bind(&q.it_product),
	)

	//选择val
	q.application.Layout(compose.Label(attr.Point(10, 45), attr.Size(20, 10), attr.Title("间隔")),
		compose.ComboBox(attr.Point(40, 40), attr.Size(160, 25), attr.Options("1m", "3m", "5m", "15m")).Bind(&q.cc_val),
		compose.CheckBox(attr.Point(240, 45), attr.Size(20, 10), attr.Title("微信Claw通知")).Bind(&q.cb_weixin),
	)

	//输入工程，开始盯盘
	q.application.Layout(compose.Label(attr.Point(10, 75), attr.Size(20, 10), attr.Title("工程")),
		compose.Input(compose.IM_Normal, attr.Point(40, 70), attr.Size(180, 25)).Bind(&q.it_project),
		compose.Button(attr.Point(240, 70), attr.Size(80, 25), attr.Title("开始")).Event(q.start))

	return q.application
}

func (q *QuantEye) start(sender compose.Component, data *compose.EventData) {
	v := q.it_product.Get()
	if q.it_product.Get() == "" {
		ui.ShowMessage("错误", "请输入盯盘的产品")
		return
	}

	if q.cc_val.Get() < 0 {
		ui.ShowMessage("错误", "请选择盯盘数据的时间间隔")
		return
	}

	if q.it_project.Get() == "" {
		ui.ShowMessage("错误", "请输入盯盘的策略工程")
		return
	}

	//启动微信claw通知
	if q.cb_weixin.Get() {
		q.weixin = &Weixin{UrlHandler: q.showUrl, Handler: q.OnUserMsg}
		q.weixin.Init()
		q.weixin.Start()
	}
	ui.ShowMessage("消息", "你点击开始了,"+v)
}

func (q *QuantEye) OnUserMsg(msg string) string {
	if strings.Contains(msg, "盯盘") || strings.Contains(msg, "做什么") {
		defaultmessage := fmt.Sprintf("正在盯盘中，产品：%s,%s线", q.it_product.Get(), q.cc_val.Get())
		return defaultmessage
	}
	return ""
}

func (q *QuantEye) showUrl(url string) {
	ui.ShowMessage("请扫码", "在浏览器中访问："+url)
}
