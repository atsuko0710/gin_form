package service

import (
	"errors"
	"gin_forum/params"
	"gin_forum/pkg/snowflake"
	"gin_forum/repository"
)

func Register(request params.CreateRequest) (err error) {

	if (!repository.CheckUserExist(request.Username)) {
		return errors.New("用户已存在")
	}

	snowflake.GetID()

	repository.CreateUser()
	return
}