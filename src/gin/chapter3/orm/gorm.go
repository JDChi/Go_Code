package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

type Product struct {
	ID     int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Number string `gorm:"unique" json:"number"`
	Name   string `gorm:"type:varchar(20);not null" json:"name"`
	MadeIn string `gorm:"type:varchar(128);not null" json:"made_in"`
}

type GormResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

var gormDB *gorm.DB
var gormResp GormResp

func init() {
	var err error
	dsn := "root:JDNEW99123@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	gormDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("connect to mysql wrong ", err)
		return
	}
}

func main() {
	r := gin.Default()
	r.POST("/gorm/insert", gormInsertData)
	r.GET("/gorm/get", gormGetData)
	r.Run(":9090")

}

func gormGetData(context *gin.Context) {
	number := context.Query("number")
	product := Product{}
	tx := gormDB.Where("number=?", number).First(&product)
	if tx.Error != nil {
		gormResp.Code = http.StatusBadRequest
		gormResp.Message = "query error"
		gormResp.Data = tx.Error
		context.JSON(http.StatusOK, gormResp)
		return
	}
	gormResp.Code = http.StatusOK
	gormResp.Message = "query succeed"
	gormResp.Data = product
	context.JSON(http.StatusOK, gormResp)
}

func gormInsertData(context *gin.Context) {
	var p Product
	err := context.Bind(&p)
	if err != nil {
		gormResp.Code = http.StatusBadRequest
		gormResp.Message = "params wrong"
		gormResp.Data = err
		context.JSON(http.StatusOK, gormResp)
		return
	}
	tx := gormDB.Create(&p)
	if tx.RowsAffected > 0 {
		gormResp.Code = http.StatusOK
		gormResp.Message = "succeed"
		gormResp.Data = "OK"
		context.JSON(http.StatusOK, gormResp)
		return
	}
}
