package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/ryanreadbooks/go-gin-examples/basic"
	"github.com/ryanreadbooks/go-gin-examples/middleware"
)

var ginEngine *gin.Engine

var (
	usebasicResponse = flag.Bool("use-basic-response", false, "Turn on APIs provided in basic response")
	usebasicRequest  = flag.Bool("use-basic-request", false, "Turn on APIs provided in basic request")
	useMiddleware    = flag.Bool("use-middleware", true, "Turn on APIs provided in middleware examples")
)

func init() {
	ginEngine = gin.Default()
}

func main() {
	flag.Parse()
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
	ginEngine.Run()
}
