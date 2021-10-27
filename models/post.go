package models

import (
	"fmt"
	"gin_forum/config/mysql"
	"gin_forum/params"
)

type Post struct {
	BaseModel
	Title      string `json:"title" gorm:"column:title;not null" bindding:"required"`
	Content    string `json:"content" gorm:"column:content;not null" bindding:"required"`
	AuthorId   int64  `json:"author_id" gorm:"column:author_id;not null" bindding:"required"`
	CategoryId int64  `json:"category_id" gorm:"column:category_id;not null" bindding:"required"`
	Status     int    `json:"status" gorm:"column:status;not null" bindding:"required"`
}

func (p *Post) TableName() string {
	return TNPost()
}

func CreatePost(p Post) (post Post, err error) {
	res := mysql.Db.Create(&p)
	return p, res.Error
}

func FindPost(Id int64) (p *params.PostDetailResponse, err error) {
	mysql.Db.Model(&Post{}).First(&p, Id)
	return
}

func FindPostWithRelate(index int64, count int64) (posts []params.ApiPostDetailResponse) {
	var p []params.ApiPostDetailResponse
	mysql.Db.Raw("SELECT `user`.`username` AS author_name,`post`.`id`,`post`.`title`,`post`.`content`,`post`.`author_id`,`post`.`status`,`post`.`category_id`,`post`.`create_time`,`category`.`name` AS `category_name` FROM `post` INNER JOIN `category` INNER JOIN `user` WHERE `post`.`category_id` = `category`.`id` AND `post`.`author_id` = `user`.`user_id` LIMIT ?, ?", (index-1)*count, count).Scan(&p)

	fmt.Println("ApiPostDetailResponse", p)
	return
}
