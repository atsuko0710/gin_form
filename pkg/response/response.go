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
	Data interface{} `json:"data,omitempty"`
}

func SendResponse(c *gin.Context, code ResCode, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: data,
	})
}

func SendResponseWithMsg(c *gin.Context, code ResCode, msg string, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}
