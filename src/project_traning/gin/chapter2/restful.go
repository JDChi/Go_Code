package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
)

// ThirdLoginReq 第三方登录请求数据
type ThirdLoginReq struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

// ThirdLoginResp 第三方返回的登录数据
type ThirdLoginResp struct {
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

// ClientReq 客户端请求数据
type ClientReq struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Other    any    `json:"other"`
}

// ClientResp 响应客户端数据
type ClientResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

// 调用 Restful 接口
func main() {
	r := gin.Default()
	r.POST("/getOtherAPI", getOtherAPI)
	r.Run(":9091")
}

func getOtherAPI(context *gin.Context) {
	// 除了客户端请求的数据
	var reqData ClientReq
	var resp ClientResp
	err := context.Bind(&reqData)
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Msg = "request params error"
		resp.Data = err
		context.JSON(http.StatusBadRequest, resp)
		return
	}
	// 请求第三方数据
	url := "http://127.0.0.1:9090/login"
	thirdLoginReq := ThirdLoginReq{reqData.UserName, reqData.Password}
	data, err := getRestfulAPI(url, thirdLoginReq, "application/json")
	var thirdLoginResp ThirdLoginResp
	json.Unmarshal(data, &thirdLoginResp)
	// 将第三方的数据转到响应客户端数据里返回
	resp.Code = http.StatusOK
	resp.Msg = "request succeed"
	resp.Data = thirdLoginResp
	context.JSON(http.StatusOK, resp)

}

func getRestfulAPI(url string, data any, contentType string) ([]byte, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println("get restful api error")
		return nil, err
	}
	res, err := ioutil.ReadAll(resp.Body)
	return res, err
}
