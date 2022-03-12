package main

import (
	"github.com/GodWY/wlog"
	"github.com/GodWY/wlog/beego"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	cc := &beego.Options{}

	for _, opt := range [...]beego.Option{
		beego.WithFilename("logs/development.log"),
		beego.WithMaxlines(10000),
		beego.WithMaxsize(1 << 28),
		beego.WithDaily(true),
		beego.WithMaxDays(7),
		beego.WithRotate(true),
		beego.WithLevel(0),
		beego.WithPerm("0600"),
		beego.WithSeparate([]string{"error"}...),
		beego.WithDebug(false),
	} {
		_ = opt(cc)
	}
	b := beego.New(cc)
	logger := wlog.NewEntry(b, "app")

	router.GET("/get", func(context *gin.Context) {
		logger.MustAppendFromGinContext(context).EventId("login").SubEventId("pppp")

		defer logger.Flush()
		logger.Info("ppppppppp")
		context.JSON(http.StatusOK, "success")
		return
	})
	//router.Use(middleware.LogForGin(logger))
	router.Run()
}
