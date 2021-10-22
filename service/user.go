package service

import (
	"errors"
	"gin_forum/models"
	"gin_forum/params"
	"gin_forum/pkg/auth"
	"gin_forum/pkg/snowflake"
	"gin_forum/repository"
)

// Register 注册用户
func Register(request params.CreateRequest) (err error) {
	if (!repository.CheckUserExist(request.Username)) {
		return errors.New("用户已存在")
	}

	userID := snowflake.GetID()
	pass, err := auth.Encrypt(request.Password)
	if err != nil {
		return errors.New("密码加密失败")
	}
	
	user := models.User{
		UserId: userID,
		UserName: request.Username,
		Password: pass,
	}

	if err :=repository.CreateUser(user); err != nil {
		return errors.New("创建用户失败")
	}
	
	return nil
}

// Login 登录逻辑
func Login(request params.LoginRequest) (err error) {
	user, err := models.FindByUsername(request.Username)
	if err != nil {
		return errors.New("该用户不存在")
	}

	if err = auth.Compare(user.Password, request.Password); err != nil {
		return errors.New("密码错误")
	}
	return nil
}