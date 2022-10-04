package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type ValidUser struct {
	Name    string  `validate:"required" ,json:"name"`
	Email   string  `validate:"email" ,json:"email"`
	Address Address `validate:"dive" ,json:"address"` // dive 关键字表示进入到嵌套结构体力进行判断
}

type Address struct {
	City  string `validate:"required" ,json:"city"`
	Phone string `validate:"numeric,len=11" ,json:"phone"`
}

var validate1 *validator.Validate

func init() {
	validate1 = validator.New()
}

// 嵌套结构体的校验
func main() {
	r := gin.Default()
	var user ValidUser
	r.POST("/validate1", func(context *gin.Context) {
		err := context.Bind(&user)
		if err != nil {
			context.JSON(http.StatusBadRequest, "request params error")
			return
		}

		if validateUser(user) {
			context.JSON(http.StatusBadRequest, "validation failed")
			return
		}
		context.JSON(http.StatusOK, "validation succeed")

	})
	r.Run(":9090")

}

func validateUser(u ValidUser) bool {
	// 校验
	err := validate1.Struct(u)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			fmt.Println("wrong field: ", e.Field())
			fmt.Println("wrong value: ", e.Value())
			fmt.Println("wrong tag: ", e.Tag())
		}
		return false
	}
	return true
}
