package params

import "time"

type CategoryListResponse struct {
	Id   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type CategoryDetailResponse struct {
	Id         int64     `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	Desc       string    `json:"desc,omitempty" db:"desc"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
}
