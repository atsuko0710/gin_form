package service

import (
	"gin_forum/models"
	"gin_forum/params"
	"gin_forum/pkg/auth"
	"gin_forum/pkg/snowflake"
	"gin_forum/pkg/response"
	"gin_forum/repository"

	"go.uber.org/zap"
)

// Register 注册用户
func Register(request params.CreateRequest) (code response.ResCode) {
	if (!repository.CheckUserExist(request.Username)) {
		return response.UserExist
	}

	userID := snowflake.GetID()
	pass, err := auth.Encrypt(request.Password)
	if err != nil {
		zap.L().Error("auth.Encrypt() failed", zap.Error(err))
		return response.InternalServerError
	}
	
	user := models.User{
		UserId: userID,
		UserName: request.Username,
		Password: pass,
	}

	if err :=repository.CreateUser(user); err != nil {
		return response.CreateUserField
	}
	
	return response.OK
}

// Login 登录逻辑
func Login(request params.LoginRequest) (resCode response.ResCode) {
	user, err := models.FindByUsername(request.Username)
	if err != nil {
		return response.UserNotExist
	}

	if err = auth.Compare(user.Password, request.Password); err != nil {
		zap.L().Error("auth.Compare() failed", zap.Error(err))
		return response.InvalidPassword
	}
	return response.OK
}