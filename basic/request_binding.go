package basic

import "github.com/gin-gonic/gin"
import "net/http"

func BindingDemo(ginEngine *gin.Engine) {
	// query bind
	type QueryStruct struct {
		Version string `form:"version" json:"version"`
		Page    int    `form:"page" json:"page"`
		Start   string `form:"start" json:"start"`
	}

	ginEngine.GET("/basic/bind-query", func(ctx *gin.Context) {
		var q QueryStruct
		var err error
		if err = ctx.ShouldBindQuery(&q); err == nil {
			// 这里就从url的query参数中构建了一个QueryStruct对象
			ctx.JSON(http.StatusOK, q)
			return
		}
		ctx.String(http.StatusOK, err.Error())
	})

	// 见在request body中的json格式数据参数绑定到结构体中（要求结构体中包含json的tag）
	ginEngine.POST("/basic/bind-json", func(ctx *gin.Context) {
		var q QueryStruct
		var err error
		if err = ctx.ShouldBindJSON(&q); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		// 此时结构体变量q就已经被绑定成功
		ctx.JSON(http.StatusOK, q)
	})

	// 由于ShouldBindQuery是用查询参数来绑定结构体，所以如果我们想要用request body中的form来绑定结构体的话
	// 需要使用ShouldBind方法
	// 结构体中的tag依然使用form
	ginEngine.POST("/basic/bind-form", func(ctx *gin.Context) {
		var q QueryStruct
		var err error
		if err = ctx.ShouldBind(&q); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		// 此时结构体变量q就已经被绑定成功
		ctx.JSON(http.StatusOK, q)
	})
}
