package models

import (
	"gin_forum/config/mysql"
	"gin_forum/params"
)

type Category struct {
	BaseModel
	Name string `json:"name" gorm:"column:name;not null" bindding:"required"`
	Desc string `json:"desc" gorm:"column:desc"`
}

func (c *Category) TableName() string {
	return TNCategory()
}

// CategoryList 获取分类列表
func CategoryList() (c []*params.CategoryListResponse, err error) {
	d := mysql.Db.Model(&Category{}).Find(&c)
	return c, d.Error
}

// FindCategory 根据ID查询数据
func FindCategory(Id int64) (c *params.CategoryDetailResponse, err error) {
	mysql.Db.Model(&Category{}).First(&c, Id)
	return
}
