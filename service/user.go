package service

import (
	"gin_forum/params"
	"gin_forum/pkg/snowflake"
	"gin_forum/repository"
)

func Register(request params.CreateRequest) (err error) {

	if (!repository.FindUserByUsername(request.Username)) {

	}

	snowflake.GetID()

	repository.CreateUser()
	return
}