package basic

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ResponseFromReader(ginEngine *gin.Engine) {
	ginEngine.GET("/basic/data-reader-local", func(ctx *gin.Context) {
		// 从reader中读取响应的内容
		// 打开文件
		file, err := os.Open("assets/plain.txt")
		defer func() { _ = file.Close() }()
		if err != nil {
			// 用ctx直接返回状态码
			// 可以将错误的内容返回
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}
		info, err := file.Stat()
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}
		ctx.DataFromReader(http.StatusOK, info.Size(), "text/plain", file, nil)
	})

	// 或者是从其它reader读入都是可以的
	ginEngine.GET("/basic/data-reader-remote", func(ctx *gin.Context) {
		// 从远程请求一张图片
		resp, err := http.Get("https://golang.google.cn/images/gophers/ladder.svg")
		if err != nil || resp.StatusCode != http.StatusOK {
			// ctx.Status(http.StatusServiceUnavailable) // 可以直接设置响应状态码
			ctx.JSON(http.StatusServiceUnavailable, gin.H{
				"msg":  err.Error(),
				"code": resp.StatusCode,
			})
			return
		}
		// 从远程请求的资源成功，可以开始返回响应结果
		contentType := resp.Header.Get("Content-Type")
		contentLength := resp.ContentLength
		contentReader := resp.Body
		defer func() { _ = resp.Body.Close() }()

		// 用reader
		ctx.DataFromReader(http.StatusOK, contentLength, contentType, contentReader, nil)
	})
}
