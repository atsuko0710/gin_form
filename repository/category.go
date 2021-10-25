package repository

import (
	"gin_forum/models"
)

func GetCategoryList() (c []*models.CategoryListApi, err error) {
	c, err = models.CategoryList()
	return
}

func GetCategoryDetail(Id int64) (c *models.CategoryDetailApi, err error)  {
	c, err = models.FindCategory(Id)
	return
}