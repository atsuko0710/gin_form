package models

import (
	"gin_forum/config/mysql"
	"time"
)

type Category struct {
	BaseModel
	Name string `json:"name" gorm:"column:name;not null" bindding:"required"`
	Desc string `json:"desc" gorm:"column:desc"`
}

type CategoryListApi struct {
	Id   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type CategoryDetailApi struct {
	Id         int64     `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	Desc       string    `json:"desc,omitempty" db:"desc"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
}

func (c *Category) TableName() string {
	return TNCategory()
}

// CategoryList 获取分类列表
func CategoryList() (c []*CategoryListApi, err error) {
	d := mysql.Db.Model(&Category{}).Find(&c)
	return c, d.Error
}

// FindCategory 根据ID查询数据
func FindCategory(Id int64) (c *CategoryDetailApi, err error) {
	mysql.Db.Model(&Category{}).First(&c, Id)
	return
}
