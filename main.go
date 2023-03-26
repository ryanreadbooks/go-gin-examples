package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/ryanreadbooks/go-gin-examples/basic"
)

var ginEngine *gin.Engine

var (
	usebasicResponse = flag.Bool("use-basic-response", true, "Turn APIs provided in basic response")
	usebasicRequest  = flag.Bool("use-basic-request", true, "Turn APIs provided in basic request")
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
	ginEngine.Run()
}
