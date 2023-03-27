package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddGroupMiddleware(ginEngine *gin.Engine) {
	g1 := ginEngine.Group("/group")
	// 在路由组中添加中间件，每个路由组之间不受影响
	g1.Use(checkUserName)

	{
		g1.GET("/login", login)
		g1.GET("/logout", logout)
	}

}

// 这个中间件做一个简单的检查工作
func checkUserName(ctx *gin.Context) {
	// 检查请求url的查询参数中是否有name，如果不存在则不继续往下走
	if name := ctx.Query("name"); name == "" {
		// name不存在或者name为空，则返回失败信息
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "denial",
			"why":     "no name param in query",
		})
		return
	}
	// 成功则可以继续往下走
	ctx.Next() // 必须显式调用
}

func login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "welcome",
		"name":    ctx.Query("name"),
	})
}

func logout(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "see you next time",
		"name":    ctx.Query("name"),
	})
}
