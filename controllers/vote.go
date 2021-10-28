package controllers

import (
	"gin_forum/params"
	"gin_forum/pkg/response"
	// "gin_forum/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// PostVote 投票
func PostVote(c *gin.Context) {
	var params params.VoteRequest
	// userId, _, _ := getCurrentUser(*c)
	if err := c.ShouldBindJSON(&params); err != nil {
		// 判断错误是不是 validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			zap.L().Error("PostVote.validator() failed", zap.Error(err))
			response.SendResponse(c, response.InvalidParam, gin.H{})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}
	// service.VoteForPost()
}
