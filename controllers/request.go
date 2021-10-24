package controllers

import (
	"gin_forum/pkg/response"
	"gin_forum/router/middleware"

	"github.com/gin-gonic/gin"
)

// getCurrentUser 获得当前登陆用户ID和用户名
func getCurrentUser(c gin.Context) (userID int64, userName string, resCode response.ResCode) {
	uid, ok := c.Get(middleware.ContextUserIDKey)
	username := c.GetString(middleware.ContextUserNameKey)
	if !ok {
		return 0, "", response.UserNotLogin
	}

	userID, ok = uid.(int64)
	if !ok {
		return 0, "", response.UserNotLogin
	}
	return userID, username, response.OK
}