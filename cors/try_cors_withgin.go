package cors

import (
	"net/http"

	ginCors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ProvideCORSAllowedWithGin(ginEngine *gin.Engine) {
	g := ginEngine.Group("/cors-gin")
	// gin的cors作为中间件来使用
	g.Use(ginCors.New(ginCors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	g.GET("/get", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	})
}
