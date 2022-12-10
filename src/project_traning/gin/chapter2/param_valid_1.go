package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"unicode/utf8"
)

// UserInfo
//go get github.com/go-playground/validator/v10
type UserInfo struct {
	Id   string `validate:"uuid" json:"id"`           // 校验 uuid 类型
	Name string `validate:"checkName" json:"name"`    // 自定义校验
	Age  uint8  `validate:"min=0,max=130" json:"age"` //
}

var validate *validator.Validate

func init() {
	validate = validator.New()
	// 注册自定义校验
	validate.RegisterValidation("checkName", checkNameFunc)

}

func checkNameFunc(fl validator.FieldLevel) bool {
	// 将名称转为 utf-8 再进行判断
	length := utf8.RuneCountInString(fl.Field().String())
	if length >= 2 && length <= 12 {
		return true
	}
	return false
}

// 参数校验 1
func main() {
	r := gin.Default()
	var user UserInfo
	r.POST("/validate", func(context *gin.Context) {
		err := context.Bind(&user)
		if err != nil {
			context.JSON(http.StatusBadRequest, "request params error")
			return
		}
		// 校验
		err = validate.Struct(user)
		if err != nil {
			for _, e := range err.(validator.ValidationErrors) {
				fmt.Println("wrong field: ", e.Field())
				fmt.Println("wrong value: ", e.Value())
				fmt.Println("wrong tag: ", e.Tag())
			}
			context.JSON(http.StatusBadRequest, "validation failed")
			return
		}

		context.JSON(http.StatusOK, "validation succeed")

	})
	r.Run(":9090")

}
