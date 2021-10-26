package repository

import (
	"gin_forum/models"
	"gin_forum/params"
)

func CreatePost(p models.Post) error {
	return models.CreatePost(p)
}

func GetPostDetail(Id int64) (c *params.PostDetailResponse, err error) {
	c, err = models.FindPost(Id)
	return
}

func GetPostList(index int64, count int64) (posts []params.ApiPostDetailResponse) {
	posts = models.FindPostWithRelate(index, count)
	return
}
