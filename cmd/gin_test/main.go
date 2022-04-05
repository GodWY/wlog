package main

import (
	"net/http"

	"github.com/GodWY/wlog"
	"github.com/GodWY/wlog/beego"
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
		beego.WithRotate(false),
		beego.WithLevel(0),
		beego.WithPerm("0600"),
		beego.WithSeparate([]string{"error", "info"}...),
		beego.WithDebug(false),
	} {
		_ = opt(cc)
	}
	xx := wlog.NewLogger(beego.New(cc), "elk")

	logger := xx.WithElkEntry("xxxx")
	router.GET("/get", func(context *gin.Context) {
		// defer logger.Flush()
		// logger.Info("ppppppppp")
		logger.MustAppendFromGinContext(context).EventId("getxxxx").UseID("1212121212121212121212")
		defer logger.Flush()
		context.JSON(http.StatusOK, "success")
	})
	//router.Use(middleware.LogForGin(logger))
	router.Run()
}
