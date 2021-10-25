package service

import (
	"database/sql"
	"gin_forum/params"
	"gin_forum/pkg/response"
	"gin_forum/repository"

	"go.uber.org/zap"
)

// GetCategoryList 获取分类列表
func GetCategoryList() (c []*params.CategoryListResponse, resCode response.ResCode) {
	c, err := repository.GetCategoryList()
	if err != nil {
		zap.L().Error("repository.GetCategoryList() failed", zap.Error(err))
		return []*params.CategoryListResponse{}, response.GetListFail
	}
	return c, response.OK
}

// GetCategoryDetail 获取分类详情
func GetCategoryDetail(Id int64) (c *params.CategoryDetailResponse, resCode response.ResCode)  {
	c, err := repository.GetCategoryDetail(Id)

	// 没有找到结果
	if err == sql.ErrNoRows {
		return &params.CategoryDetailResponse{}, response.GetDetailFail
	}
	
	if err != nil {
		zap.L().Error("repository.GetCategoryDetail() failed", zap.Error(err))
		return &params.CategoryDetailResponse{}, response.GetDetailFail
	}
	return c, response.OK
}