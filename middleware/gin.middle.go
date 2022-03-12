package middleware

import (
	"github.com/gin-gonic/gin"
	"tygit.tuyoo.com/gocomponents/tylog"
)

// MetricForGin 基于gin的日志中间件。
func MetricForGin() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 读取请求的参数

		return
	}
}

// MetricForRpc rpc日志
func MetricForRpc() {

}

// GaForRpc ga日志
func GaForRpc() {

}

// GaForGin ga日志for gin
func GaForGin() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

// LogForGin 普通日志
func LogForGin(log *tylog.Entry) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 开始日志打印
		//	trace:      ctx.Get("trace").(log.TraceSpan),
		//		userAgent:  ctx.Request().Header.Get("User-Agent"),
		//		ip:         ctx.RealIP(),
		//		path:       ctx.Request().URL.Path,
		//		method:     ctx.Request().Method,
		//		proto:      ctx.Request().Proto,
		//		status:     ctx.Response().Status,
		//		accessId:   ctx.Request().Header.Get("Access-Id"),
		//		signType:   ctx.Request().Header.Get("Sign-Type"),
		//		secretType: ctx.Request().Header.Get("Secret-Type"),
		//		random:     ctx.Request().Header.Get("Random"),
		//		timestamp:  ctx.Request().Header.Get("Timestamp"),
		//		signature:  ctx.Request().Header.Get("Signature"),
		log.MustAppend().String("user-agent", ctx.Request.Header.Get("User-Agent")).
			String("ip", ctx.ClientIP()).String("path", ctx.Request.URL.Path).
			String("method", ctx.Request.Method).String("proto", ctx.Request.Proto)

		defer log.Flush()
		//ctx.Next()
		log.MustAppend().String("resp", "ssss")
		return
	}
}

func LogForRpc() {
}
