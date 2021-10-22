package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
返回结构
{
	"code": 100,
	"msg": "success"
	"data": {},
}
*/

type ResponseData struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func SendResponse(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

func SendResponseWithMsg(c *gin.Context, code ResCode, msg string) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}
