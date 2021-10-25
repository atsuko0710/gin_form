package params

import "time"

type CreatePostRequest struct {
	Title      string `json:"title" bindding:"required"`
	Content    string `json:"content" bindding:"required"`
	AuthorId   int64  `json:"author_id" bindding:"required"`
	Status     string `json:"status" bindding:"required"`
	CategoryId string `json:"category_id" bindding:"required"`
}

type PostDetailResponse struct {
	Id         int64     `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	AuthorId   int64     `json:"author_id"`
	Status     string    `json:"status"`
	CategoryId string    `json:"category_id"`
	CreateTime time.Time `json:"create_time"`
}
