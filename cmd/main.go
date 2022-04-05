package main

import (
	"time"

	"tygit.tuyoo.com/gocomponents/tylog"
	"tygit.tuyoo.com/gocomponents/tylog/beego"
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
	xx := tylog.NewLogger(beego.New(cc), "myself")
	for i := 0; i < 100; i++ {
		go func() {
			logger := xx.WithEntry("findId")
			logger.MustAppend().String("Xxxxx", "1111111").Int32("aaaa", 111)

			logger.Info("xxxx121212")

			logger.MustAppend().String("yyyyyy", "1111111").Int32("bbbb", 222)

			logger.Error("abc")
			logger.Debug("ppp")
			// xxxx121212
			logger.Flush()
		}()
	}
	time.Sleep(5 * time.Second)
}
