package service

import (
	"errors"
	"gin_forum/models"
	"gin_forum/params"
	"gin_forum/pkg/snowflake"
	"gin_forum/repository"
)

// Register 注册用户
func Register(request params.CreateRequest) (err error) {
	if (!repository.CheckUserExist(request.Username)) {
		return errors.New("用户已存在")
	}

	userID := snowflake.GetID()
	user := models.User{
		UserId: userID,
		UserName: request.Username,
		Password: request.Password,
	}

	if err :=repository.CreateUser(user); err != nil {
		return errors.New("创建用户失败")
	}
	
	return nil
}