package service

import (
	"gin_forum/models"
	"gin_forum/params"
	"gin_forum/pkg/response"
	"gin_forum/repository"
	"strconv"
)

const POST_NORMAL_STATUS = 1

func CreatePost(request params.CreatePostRequest) response.ResCode {
	categoryId, err := strconv.ParseInt(request.CategoryId, 10, 64)
	if err != nil {
		return response.InvalidParam
	}
	
	post := models.Post{
		Title: request.Title,
		Content: request.Content,
		CategoryId: categoryId,
		AuthorId: request.AuthorId,
		Status: POST_NORMAL_STATUS,
	}

	if err := repository.CreatePost(post); err != nil {
		return response.CreatePostFail
	}
	return response.OK
}
