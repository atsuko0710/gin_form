package params

import "time"

type CategoryListResponse struct {
	Id   int64  `json:"id" bindding:"required"`
	Name string `json:"name" bindding:"required"`
}

type CategoryDetailResponse struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`
	Desc       string    `json:"desc,omitempty"`
	CreateTime time.Time `json:"create_time"`
}
