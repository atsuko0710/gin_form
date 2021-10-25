package models

import (
	"gin_forum/config/mysql"
	"gin_forum/params"
)

// import "gin_forum/config/mysql"

type Post struct {
	BaseModel
	Title    string `json:"title" gorm:"column:title;not null" bindding:"required"`
	Content  string `json:"content" gorm:"column:content;not null" bindding:"required"`
	AuthorId int64  `json:"author_id" gorm:"column:author_id;not null" bindding:"required"`
	CategoryId int64  `json:"category_id" gorm:"column:category_id;not null" bindding:"required"`
	Status   int    `json:"status" gorm:"column:status;not null" bindding:"required"`
}

func (p *Post) TableName() string {
	return TNPost()
}

func CreatePost(p Post) error {
	return mysql.Db.Create(&p).Error
}

func FindPost(Id int64) (p *params.PostDetailResponse, err error) {
	mysql.Db.Model(&Post{}).First(&p, Id)
	return
}