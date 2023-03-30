package main

import (
	"flag"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ryanreadbooks/go-gin-examples/basic"
	"github.com/ryanreadbooks/go-gin-examples/cors"
	"github.com/ryanreadbooks/go-gin-examples/middleware"
)

var ginEngine *gin.Engine

var (
	usebasicResponse = flag.Bool("use-basic-response", false, "Turn on APIs provided in basic response")
	usebasicRequest  = flag.Bool("use-basic-request", false, "Turn on APIs provided in basic request")
	useMiddleware    = flag.Bool("use-middleware", true, "Turn on APIs provided in middleware examples")
	useCustomLog     = flag.Bool("use-custom-log", false, "Turn on examples using custom logs")
	useCors          = flag.Bool("use-cors", true, "Turn on examples using cors")
)

func init() {
	ginEngine = gin.Default()
}

func main() {
	flag.Parse()

	ginEngine.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  "/ ok",
			"code": 200,
		})
	})

	if *usebasicResponse {
		basic.ProvideBasicContextResponse(ginEngine)
	}
	if *usebasicRequest {
		basic.ProvideBasicContextRequest(ginEngine)
	}
	if *useMiddleware {
		middleware.AddGlobalMiddleware(ginEngine)
		middleware.AddGroupMiddleware(ginEngine)
		middleware.MiddleCommunication(ginEngine)
	}
	if *useCustomLog {
		gin.DisableConsoleColor() // 取消控制台日志的颜色
		// 将日志写入控制台的同时也写入文件
		f, _ := os.OpenFile("log", os.O_RDWR|os.O_CREATE, 0664)
		gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	}
	if *useCors {
		// 再跑一个服务用来提供cors.html文件的服务
		go func() {
			frontEndServer := gin.Default()
			frontEndServer.StaticFile("/cors.html", "./assets/cors.html")
			frontEndServer.Run(":8081")
		}()
		cors.ProvideCORSAllowedWithGin(ginEngine)
		cors.ProvideCORSAllowed(ginEngine)
	}
	ginEngine.Run(":8080")
}
