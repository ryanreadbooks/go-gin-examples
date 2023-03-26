package basic

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// gin.Context中和响应相关的方法怎样使用
// gin.Context中包含了许多用来返回响应的方法
func ProvideBasicContextResponse(ginEngine *gin.Engine) {

	ginEngine.GET("/basic/ascii-json", func(ctx *gin.Context) {
		res := struct {
			Code    int
			Message string
			Version string
			People  struct {
				Name string
				Age  int
			}
		}{
			200,
			"OK",
			"v1.0.0",
			struct {
				Name string
				Age  int
			}{"张三三", 19}, // 这里是匿名的结构体嵌套，所以必须要指明类型，不然报错
		}

		// 返回response
		ctx.AsciiJSON(http.StatusOK, res) // 以ascii格式返回响应，如果当中有非ascii的内容，则进行自动转义操作
	})

	ginEngine.GET("/basic/json", func(ctx *gin.Context) {
		res := struct {
			Code    int
			Message string
			Version string
			People  struct {
				Name string
				Age  int
			}
		}{
			200,
			"OK",
			"v1.0.0",
			struct {
				Name string
				Age  int
			}{"张三三", 19}, // 这里是匿名的结构体嵌套，所以必须要指明类型，不然报错
		}

		// 返回response
		ctx.JSON(http.StatusOK, res) // 以utf-8编码形式返回响应
	})

	// 返回纯文本
	ginEngine.GET("/basic/plain-text", func(ctx *gin.Context) {
		ctx.Data(http.StatusOK, "text/plain; charset=utf-8", []byte("Hello From Gin, 来自Gin的问候"))
	})

	// 返回字符串，可以用format的形式组织字符串
	ginEngine.GET("/basic/string", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "%s:%s:%d", "hello", ctx.ClientIP(), 100)
	})

	// 也可以直接响应文件
	ginEngine.GET("/basic/html", func(ctx *gin.Context) {
		ctx.File("assets/test.html")
	})

	FileAttachmentDemo(ginEngine)  // 文件下载
	ResponseFromReader(ginEngine)  // 从Reader中读入响应内容
	ResponseRedirecDemo(ginEngine) // 重定向

}
