package middleware

import (
	"github.com/gin-gonic/gin"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func MiddleCommunication(ginEngine *gin.Engine) {
	g1 := ginEngine.Group("/comm")
	g1.Use(middleware1)
	{
		g1.GET("/ping", func(ctx *gin.Context) {
			// 可以从Context中取出key-value
			if p, ok := ctx.Get("ryan"); ok {
				pp := p.(*Person)
				ctx.JSON(200, pp)
				return
			}
			ctx.JSON(200, gin.H{
				"msg": "error",
			})
		})
	}
}

func middleware1(c *gin.Context) {
	c.Set("ryan", &Person{Name: "ryan", Age: 19}) // 在Context中存入一个值
}
