package service

import (
	"database/sql"
	"gin_forum/models"
	"gin_forum/params"
	"gin_forum/pkg/response"
	"gin_forum/repository"
	"strconv"

	"go.uber.org/zap"
)

const POST_NORMAL_STATUS = 1

func CreatePost(request params.CreatePostRequest) response.ResCode {
	categoryId, err := strconv.ParseInt(request.CategoryId, 10, 64)
	if err != nil {
		return response.InvalidParam
	}

	post := models.Post{
		Title:      request.Title,
		Content:    request.Content,
		CategoryId: categoryId,
		AuthorId:   request.AuthorId,
		Status:     POST_NORMAL_STATUS,
	}

	if err := repository.CreatePost(post); err != nil {
		return response.CreatePostFail
	}
	return response.OK
}

func GetPostDetail(Id int64) (c *params.PostDetailResponse, resCode response.ResCode) {
	c, err := repository.GetPostDetail(Id)

	// 没有找到结果
	if err == sql.ErrNoRows {
		return &params.PostDetailResponse{}, response.GetDetailFail
	}

	if err != nil {
		zap.L().Error("repository.GetPostDetail() failed", zap.Error(err))
		return &params.PostDetailResponse{}, response.GetDetailFail
	}
	return c, response.OK
}

func GetPostList(index int64, count int64) (posts []params.ApiPostDetailResponse) {
	posts = repository.GetPostList(index, count)
	return
}
