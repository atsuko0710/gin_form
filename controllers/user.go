package controllers

import (
	"net/http"

	"gin_forum/params"
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
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "参数有误",
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}
	
	if err := service.Register(params); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

// Login 登录入口
func Login(c *gin.Context)  {
	// 获取参数
	var params params.LoginRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		// 判断错误是不是 validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "参数有误",
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}

	if err := service.Login(params); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
