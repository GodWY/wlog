package main

import (
	"github.com/GodWY/wlog"
	"github.com/GodWY/wlog/beego"
)

func main() {
	// 第一种 函数执行完 统一打印
	cc := &beego.Options{}
	for _, opt := range [...]beego.Option{
		beego.WithFilename("logs/development.log"),
		beego.WithMaxlines(10000),
		beego.WithMaxsize(1 << 28),
		beego.WithDaily(true),
		beego.WithMaxDays(7),
		beego.WithRotate(false),
		beego.WithLevel(0),
		beego.WithPerm("0600"),
		beego.WithSeparate([]string{"error", "info"}...),
		beego.WithDebug(false),
	} {
		_ = opt(cc)
	}

	logger := wlog.NewEntry(beego.New(cc), "msg")
	logger.MustAppend().EventId("xxxx").String("Xxxxx", "1111111")
	logger.Flush()
	logger.Info("xxxx121212")

	// 第二种 直接打印
	logger2 := wlog.NewEntry(beego.New(cc), "msg")
	logger2.MustAppend().Info("xxxxxxxxx")

}
