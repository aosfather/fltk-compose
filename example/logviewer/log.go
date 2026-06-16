package main

import (
	"time"

	"github.com/aosfather/fltk-compose/attr"
	"github.com/aosfather/fltk-compose/compose"
)

func main() {
	var log compose.BindObj[compose.Log]
	app := compose.App(attr.Size(400, 500),
		attr.Title("盯盘"))
	app.Layout(compose.LogView(attr.Point(10, 10), attr.Size(100, 200)).Bind(&log))
	go func() {
		time.Sleep(3 * time.Second)
		l := log.Get()
		l.Debug("haha")
		l.Info("yes")
		l.Warn("sorry")
		l.Error("danger")
		l.ToTop()
	}()
	app.Run()

}
