package basic

import "github.com/gin-gonic/gin"

func FileAttachmentDemo(ginEngine *gin.Engine) {
	// 提供给用户下载文件的响应
	ginEngine.GET("/basic/html-attachment", func(ctx *gin.Context) {
		ctx.FileAttachment("assets/test.html", "cool.html")
	})
}