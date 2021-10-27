package service

import (
	"database/sql"
	"fmt"
	"gin_forum/models"
	"gin_forum/params"
	"gin_forum/pkg/enum"
	"gin_forum/pkg/response"
	"gin_forum/pkg/util"
	"gin_forum/repository"
	"strconv"
	"time"

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

	if res, err := repository.CreatePost(post); err != nil {
		return response.CreatePostFail
	}

	return response.OK
}

func createPostRedis(post models.Post) {
	now := float64(time.Now().Unix())
	voteKey := enum.KeyPostVotedPrefix + fmt.Sprint(post.Id)

	postInfo := map[string]interface{}{
		"title":    post.Title,
		"summary":  util.TruncateByWords(post.Content, 120),
		"post:id":  post.Id,
		"user:id":  post.AuthorId,
		"time":     now,
		"votes":    1,
		"comments": 0,
	}

	// 事务操作
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
