package middleware

import (
	"fmt"
	"gin_forum/pkg/response"
	"gin_forum/pkg/token"

	"github.com/gin-gonic/gin"
)

const (
	ContextUserIDKey = "userID"
	ContextUserNameKey = "userName"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.SendResponseWithMsg(c, response.InvalidToken, "请求头缺少Auth Token", gin.H{})
			c.Abort()
			return
		}
		var t string
		fmt.Sscanf(authHeader, "Bearer %s", &t)
		mc, resCode := token.ParseToken(t)
		if resCode != response.OK {
			response.SendResponse(c, resCode, gin.H{})
			c.Abort()
			return
		}
		c.Set(ContextUserIDKey, mc.UserId)
		c.Set(ContextUserNameKey, mc.Username)
		c.Next()
	}
}
