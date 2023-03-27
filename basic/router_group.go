package basic

import (
	"github.com/gin-gonic/gin"
)

func RouterGroupDemo(ginEngine *gin.Engine) {
	// gin中的路由可以分组，使它们有相同的前缀，因此可以对api进行分组管理
	// 也可以分组加中间件

	// 返回的类型是指针
	var v1 *gin.RouterGroup = ginEngine.Group("/basic/v1")
	// 通常习惯使用{}来分割每组路由
	{
		v1.GET("/login", loginV1)
		v1.GET("/info", infoV1)
	}

	v2 := ginEngine.Group("/basic/v2")
	{
		v2.GET("/login", loginV2)
		v2.GET("/info", infoV2)
		// 路由组支持嵌套
		testing := v2.Group("/testing")
		testing.GET("/do", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "testing/do",
			})
		})
	}

}

func loginV1(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "login ok",
		"version": "v1",
	})
}

func infoV1(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "info ok",
		"version": "v1",
	})
}

func loginV2(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "login ok",
		"version": "v2",
	})
}

func infoV2(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "info ok",
		"version": "v2",
	})
}