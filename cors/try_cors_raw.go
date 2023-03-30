package cors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProvideCORSAllowed(ginEngine *gin.Engine) {
	corsGroup := ginEngine.Group("/cors")
	corsGroup.GET("/get", func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST")
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "记得一键三连",
		})
	})
}