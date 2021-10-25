package repository

import (
	"gin_forum/models"
	"gin_forum/params"
)

func GetCategoryList() (c []*params.CategoryListResponse, err error) {
	c, err = models.CategoryList()
	return
}

func GetCategoryDetail(Id int64) (c *params.CategoryDetailResponse, err error)  {
	c, err = models.FindCategory(Id)
	return
}