package Middlewares

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

const (
	SuccessCode      ResponseCode = 200   //请求成功
	TokenExpiredCode ResponseCode = 40101 //token过期
	RequestFail      ResponseCode = 50001 //请求失败
)

type ResponseCode int

type Response struct {
	ErrorCode ResponseCode `json:"errno"`
	ErrorMsg  string       `json:"errmsg"`
	Data      interface{}  `json:"data"`
}

func ResponseError(c *gin.Context, code ResponseCode, err error) {
	resp := &Response{ErrorCode: code, ErrorMsg: err.Error(), Data: ""}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
}

func ResponseSuccess(c *gin.Context, data interface{}, msg string) {
	resp := &Response{ErrorCode: SuccessCode, ErrorMsg: msg, Data: data}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
}
