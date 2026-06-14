package main

import (
	"fmt"
	"strings"
	"time"

	"quanteye_utils/cli"
	"quanteye_utils/jsons"

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
	cc_valText  compose.BindObj[string]
	it_project  compose.BindObj[string]
	cb_weixin   compose.BindObj[bool]
	bt_start    compose.BindObj[bool]
	messages    compose.BindObj[compose.MessageList]
	weixin      *Weixin
	recentValue *QuantEyeRecent
	canNotify   bool
}

func (q *QuantEye) Layout(m ...compose.Modifier) *compose.Application {
	q.application = compose.App(m...)
	q.application.Layout(compose.Label(attr.Point(10, 15), attr.Size(20, 10), attr.Title("产品")),
		compose.Input(compose.IM_Normal, attr.Point(40, 10), attr.Size(180, 25)).Bind(&q.it_product),
	)

	//选择val
	q.application.Layout(compose.Label(attr.Point(10, 45), attr.Size(20, 10), attr.Title("间隔")),
		compose.ComboBox(attr.Point(40, 40), attr.Size(160, 25), attr.Options("1m", "3m", "5m", "15m")).Bind(&q.cc_val).BindText(&q.cc_valText),
		compose.CheckBox(attr.Point(240, 45), attr.Size(20, 10), attr.Title("微信Claw通知")).Bind(&q.cb_weixin),
	)

	//输入工程，开始盯盘
	q.application.Layout(compose.Label(attr.Point(10, 75), attr.Size(20, 10), attr.Title("工程")),
		compose.Input(compose.IM_Normal, attr.Point(40, 70), attr.Size(180, 25)).Bind(&q.it_project),
		compose.Button(attr.Point(240, 70), attr.Size(80, 25), attr.Title("开始")).Event(q.start).Bind(&q.bt_start))

	//消息
	q.application.Layout(compose.Messages(attr.Point(10, 100), attr.Size(380, 390), attr.Widths(380)).Bind(&q.messages))
	return q.application
}

func (q *QuantEye) start(sender compose.Component, data *compose.EventData) {
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
		q.weixin.Start()
	}

	q.bt_start.Set(false)
	//保存最新的选择
	if q.recentValue != nil {
		q.recentValue.Product = q.it_product.Get()
		q.recentValue.Val = q.cc_val.Get()
		q.recentValue.Project = q.it_project.Get()
		q.recentValue.Save()
	}
	//处理project名称
	project := q.it_project.Get()
	if !strings.HasSuffix(project, ".eye") {
		project += ".eye"
	}

	//cmd命令执行20秒后再处理通知，用于过滤历史消息
	time.AfterFunc(20*time.Second, func() { q.canNotify = true })

	//构建cmd命令，并执行
	path := cli.GetAppPath("./")
	app := path + "eombot"
	cmd := cli.NewCmdSource(app, q.showOut, "")
	cmd.Open("--source=real", "--target="+q.it_product.Get(), "--val="+q.cc_valText.Get(), "eye", project)
	cmd.Open("pwd")
}

func (q *QuantEye) OnUserMsg(msg string) string {
	if strings.Contains(msg, "盯盘") || strings.Contains(msg, "做什么") {
		defaultmessage := fmt.Sprintf("正在盯盘中，产品：%s,%d线", q.it_product.Get(), q.cc_val.Get())
		return defaultmessage
	}
	return ""
}

func (q *QuantEye) showUrl(url string) {
	ui.ShowMessage("请扫码", "在浏览器中访问："+url)
}

func (q *QuantEye) AfterRender(sender compose.Component, data *compose.EventData) {
	q.recentValue = &QuantEyeRecent{}
	q.recentValue.Init(_File, q.recentValue)
	q.recentValue.Load()
	q.it_product.Set(q.recentValue.Product)
	q.it_project.Set(q.recentValue.Project)
	q.cc_val.Set(q.recentValue.Val)
}

func (q *QuantEye) showOut(text string) {
	lay := "%s"
	if strings.Contains(text, "sell") {
		lay = "@C10@m%s"
		// say("卖出信号")
	} else if strings.Contains(text, "buy") {
		lay = "@C1@m%s"
		// say("买入信号")
	}

	ml := q.messages.Get()
	ml.AddString(lay, text)
	// q.notify(text)
	ml.ToEnd()

}

const _File = "./last.json"

type QuantEyeRecent struct {
	jsons.Recent
	Product string `json:"product"`
	Val     int    `json:"val"`
	Project string `json:"project"`
}
