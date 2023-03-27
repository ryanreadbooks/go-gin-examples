package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// 添加全局的中间件
func AddGlobalMiddleware(ginEngine *gin.Engine) {
	ginEngine.Use(tick)
}

// 这个全局中间件的功能为计算所有的耗时
func tick(c *gin.Context) {

	start := time.Now().Nanosecond()

	// gin.Context.Next直接去调用后续的中间件或者handler处理函数
	c.Next()

	// 得到耗时时间
	end := time.Now().Nanosecond()
	log.Printf("Cost %v nanoseconds\n", end-start)

}
