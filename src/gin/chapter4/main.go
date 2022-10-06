package main

import (
	"Go_Code/src/gin/chapter4/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func initConfig() {
	dir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	//viper.AddConfigPath(dir + "/config") // 执行 go run 对应的路径
	viper.AddConfigPath(dir + "/src/gin/chapter4/config/") // 执行单文件运行
	fmt.Println(dir)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	initConfig()
	common.InitDB()
	r := gin.Default()
	port := viper.GetString("server.port")
	r.Run(":" + port)

}
