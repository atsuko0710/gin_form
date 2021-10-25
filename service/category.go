package service

import (
	"gin_forum/models"
	"gin_forum/pkg/response"
	"gin_forum/repository"

	"go.uber.org/zap"
)

// GetCategoryList 获取分类列表
func GetCategoryList() (c []*models.CategoryListApi, resCode response.ResCode) {
	c, err := repository.GetCategoryList()
	if err != nil {
		zap.L().Error("repository.GetCategoryList() failed", zap.Error(err))
		return []*models.CategoryListApi{}, response.GetListFail
	}
	return c, response.OK
}

// GetCategoryDetail 获取分类详情
func GetCategoryDetail(Id int64) (c *models.CategoryDetailApi, resCode response.ResCode)  {
	c, err := repository.GetCategoryDetail(Id)
	if err != nil {
		zap.L().Error("repository.GetCategoryDetail() failed", zap.Error(err))
		return &models.CategoryDetailApi{}, response.GetDetailFail
	}
	return c, response.OK
}