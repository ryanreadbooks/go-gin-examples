package basic

import (
	"net/http"
	"reflect"
	"fmt"
	
	"github.com/go-playground/validator/v10"
	"github.com/gin-gonic/gin"
)

type Person struct {
	Name string `json:"name" binding:"required,min=2,max=10" msg:"姓名不能为空且长度只能为2~10之间"` // 这里的一个中文字算一个字符
	Age  int    `json:"age" binding:"required,gt=0,lt=150" msg:"年龄不能为空且必须在0~150范围"`
}

func BindingValidateDemo(ginEngine *gin.Engine) {
	// 模型绑定和验证

	ginEngine.POST("/basic/validate-person", func(ctx *gin.Context) {
		var p Person
		obj := reflect.TypeOf(p)
		// err的具体类型是validator.ValidationErrors
		var msgs []string = make([]string, 0, 2)
		if err := ctx.ShouldBindJSON(&p); err != nil {
			var errs validator.ValidationErrors = err.(validator.ValidationErrors)
			for _, e := range errs {
				var errField string = e.Field() // 出错的Person中字段名，比如是Name或者Age
				// 通过反射获取Person对应字段名中的Tag
				if field, ok := obj.FieldByName(errField); ok {
					msgs = append(msgs, field.Tag.Get("msg"))
				}
			}

			ctx.JSON(http.StatusOK, gin.H{
				"msg":  msgs,
				"type": fmt.Sprintf("%T", err),
			})
			return
		}
		ctx.JSON(http.StatusOK, p)
	})
}
