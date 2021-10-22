package controllers

import (
	"net/http"

	"gin_forum/params"
	"gin_forum/pkg/response"
	"gin_forum/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Register 处理用户注册入口
func Register(c *gin.Context) {
	// 获取参数
	var params params.CreateRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		// 判断错误是不是 validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.SendResponse(c, response.InvalidParam, gin.H{})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}
	resCode := service.Register(params)
	response.SendResponse(c, resCode, gin.H{})
}

// Login 登录入口
func Login(c *gin.Context)  {
	// 获取参数
	var params params.LoginRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		// 判断错误是不是 validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.SendResponse(c, response.InvalidParam, gin.H{})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}

	loginResponse, resCode := service.Login(params)
	if resCode == response.OK {
		response.SendResponse(c, resCode, gin.H{
			"accessToken":  loginResponse.AccessToken,
			"refreshToken": loginResponse.RefreshToken,
			"userID":       loginResponse.UserId,
			"username":     loginResponse.Username,
		})
	} else {
		response.SendResponse(c, resCode, gin.H{})
	}
	
}
