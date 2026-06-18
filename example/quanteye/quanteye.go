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
	app.Layout(attr.Size(400, 500), attr.Title("QuantEye-亮眼")).Run()
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
	//回测
	product    compose.BindObj[string]
	project    compose.BindObj[string]
	view       compose.BindObj[compose.Log]
	progress   compose.BindObj[float64]
	back_start compose.BindObj[bool]

	weixin      *Weixin
	recentValue *QuantEyeRecent
	canNotify   bool
	sellFormat  string
	buyFormat   string
}

func (q *QuantEye) Layout(m ...compose.Modifier) *compose.Application {
	q.application = compose.App(m...)
	q.tabs()
	q.application.Event(q.AfterRender)
	return q.application
}

func (q *QuantEye) tabs() {
	tabs := compose.Tabs(attr.M().Point(0, 0).Size(400, 500).M()...)
	tabs.NewTab("盯盘",
		compose.Label(attr.M().Point(10, 15).Size(20, 10).Title("产品").M()...),
		compose.Input(compose.IM_Normal, attr.M().Point(40, 10).Size(180, 25).M()...).Bind(&q.it_product),
		compose.Label(attr.M().Point(10, 45).Size(20, 10).Title("间隔").M()...),
		compose.ComboBox(attr.M().Point(40, 40).Size(160, 25).Options("1m", "3m", "5m", "15m").M()...).Bind(&q.cc_val).BindText(&q.cc_valText),
		compose.CheckBox(attr.M().Point(240, 45).Size(20, 10).Title("微信Claw通知").M()...).Bind(&q.cb_weixin),
		compose.Label(attr.M().Point(10, 75).Size(20, 10).Title("工程").M()...),
		compose.Input(compose.IM_Normal, attr.M().Point(40, 70).Size(180, 25).M()...).Bind(&q.it_project),
		compose.Button(attr.M().Point(240, 70).Size(80, 25).Title("开始").M()...).Event(q.start).Bind(&q.bt_start),
		compose.Messages(attr.M().Point(10, 100).Size(380, 375).Widths(380).M()...).Bind(&q.messages))

	tabs.NewTab("回测",
		compose.Label(attr.M().Point(10, 15).Size(20, 10).Title("产品").M()...),
		compose.Input(compose.IM_Normal, attr.M().Point(40, 10).Size(180, 25).M()...).Bind(&q.product),
		compose.Label(attr.M().Point(10, 45).Size(20, 10).Title("工程").M()...),
		compose.Input(compose.IM_Normal, attr.M().Point(40, 40).Size(180, 25).M()...).Bind(&q.project),
		compose.Button(attr.M().Point(240, 40).Size(80, 25).Title("开始回测").M()...).Event(q.startBackTest).Bind(&q.back_start),
		compose.LogView(attr.M().Point(10, 70).Size(380, 405).M()...).Bind(&q.view))

	q.application.Layout(tabs)
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
	if q.sellFormat == "" {
		q.sellFormat = compose.NewStyle().Color(10).Medium().ToFormat() + "%s"
	}
	if q.buyFormat == "" {
		q.buyFormat = compose.NewStyle().Color(1).Medium().ToFormat() + "%s"
	}
	q.showOut("欢迎使用QuantEye！")
	q.showOut("点击 ‘开始’  按钮")
}

func (q *QuantEye) showOut(text string) {
	lay := "%s"
	if strings.Contains(text, "sell") {
		lay = q.sellFormat
		// say("卖出信号")
	} else if strings.Contains(text, "buy") {
		lay = q.buyFormat
		// say("买入信号")
	}

	ml := q.messages.Get()
	ml.AddString(lay, text)
	q.notify(text)
	ml.ToEnd()

}

func (q *QuantEye) notify(msg string) {
	if q.canNotify && q.weixin != nil {
		q.weixin.Notify(msg)
	}
}

func (q *QuantEye) startBackTest(sender compose.Component, data *compose.EventData) {

	product := q.product.Get()
	project := q.project.Get()
	if product == "" || project == "" {
		ui.ShowMessage("错误", "请输入产品和需要回测的策略工程")
		return
	}

	q.back_start.Set(false)
	//构建cmd命令，并执行
	path := cli.GetAppPath("./")
	app := path + "eombot"
	cmd := cli.NewCmdSource(app, q.showBackTestOut, "")
	cmd.Open("--source=real", "--target="+product, "back", project)
}

func (q *QuantEye) showBackTestOut(text string) {
	log := q.view.Get()
	if strings.Contains(text, "sell") {
		log.Warn(text)
	} else if strings.Contains(text, "buy") {
		log.Debug(text)
	}

}

const _File = "./last.json"

type QuantEyeRecent struct {
	jsons.Recent
	Product string `json:"product"`
	Val     int    `json:"val"`
	Project string `json:"project"`
}
