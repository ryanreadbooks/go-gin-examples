package basic

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 学习gin.Context的和请求相关的方法
func ProvideBasicContextRequest(ginEngine *gin.Engine) {
	// 获取请求中的content-type
	ginEngine.POST("/basic/content-type", func(ctx *gin.Context) {
		var contentType string = ctx.ContentType()
		if contentType == "" {
			contentType = "empty"
		}
		ctx.String(http.StatusOK, "The content type of your request is %s", contentType)
	})

	// 请求url上的查询参数：query
	ginEngine.GET("/basic/query", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, ctx.Request.URL.Query()) // ctx.Request是net/http中的类型
	})

	// 获取post请求中的请求表单
	ginEngine.POST("/basic/post-form", func(ctx *gin.Context) {
		ctx.Request.ParseForm() // 本质上还是net/http中的
		ctx.JSON(http.StatusOK, ctx.Request.PostForm)
	})

	// 路径参数的获取
	ginEngine.GET("/basic/param/:id", func(ctx *gin.Context) {
		idVal := ctx.Param("id") // 获取URL上:id位置的值
		ctx.String(http.StatusOK, "路径参数id的值为: %s", idVal)
	})

	BindingDemo(ginEngine)         // 模型参数绑定
	BindingValidateDemo(ginEngine) // 模型参数绑定和简单校验
	UploadFileDemo(ginEngine)      // 文件上传功能

}
