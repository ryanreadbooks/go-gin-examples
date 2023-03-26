package basic

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseRedirecDemo(ginEngine *gin.Engine) {
	// 重定向到本服务器内的资源,也叫作路由重定向
	ginEngine.GET("/basic/redirect/:newlocation", func (ctx *gin.Context) {
		newlocation := ctx.Param("newlocation")
		if newlocation == "" {
			newlocation = "plain-text"
		}
		ctx.Request.URL.Path = fmt.Sprintf("/basic/%s", newlocation)
		ginEngine.HandleContext(ctx)	// 路由重定向
	})

	// 重定向到一个指定的地址,这个loc查询参数要指定完整的http url
	ginEngine.GET("/basic/redirect-to", func(ctx *gin.Context) {
		newLoc := ctx.Query("loc")
		if newLoc == "" {
			newLoc = "https://pkg.go.dev/github.com/gin-gonic/gin/"
		}
		ctx.Redirect(http.StatusMovedPermanently, newLoc + "/")
	})
}