package basic

import (
	"byex.io/irisrv/utils"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"time"
)

func Before(ctx iris.Context) {
	remoteAddr := utils.GetRemoteAddr(ctx)
	ctx.Values().Set("RemoteAddr", remoteAddr)
	println("before ########### 00000")
	ctx.Next()
}

func After(ctx iris.Context) {
	println("after  ########### 11111")
}

func customerLogger() iris.Handler {
	config := logger.Config{
		// Status displays status code
		Status: true,
		// IP displays request's remote address
		IP: true,
		// Method displays the http method
		Method: true,
		// Path displays the request path
		Path: true,
		// Query appends the url query to the Path.
		Query: true,

		//Columns: true,

		// if !empty then its contents derives from `ctx.Values().Get("logger_message")
		// will be added to the logs.
		MessageContextKeys: []string{"logger_message"},

		// if !empty then its contents derives from `ctx.GetHeader("User-Agent")
		MessageHeaderKeys: []string{"User-Agent"},

		LogFunc: func(now time.Time, latency time.Duration, status, ip, method, path string, message interface{}, headerMessage interface{}) {
			output := logger.Columnize(now.Format("2006/01/02 - 15:04:05"), latency, status, ip, method, path, message, headerMessage)
			println(output)
		},
	}

	return logger.New(config)
}
