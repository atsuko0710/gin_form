package controllers

import (
	"gin_forum/pkg/response"
	"gin_forum/service"

	"github.com/gin-gonic/gin"
	"strconv"
)

func CategoryList(c *gin.Context) {
	cateList, resCode := service.GetCategoryList()
	response.SendResponse(c, resCode, cateList)
}

func CategoryDetail(c *gin.Context) {
	idStr := c.Param("id")
	CategoryId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.SendResponse(c, response.InvalidParam, gin.H{})
	}
	cateDetail, resCode := service.GetCategoryDetail(CategoryId)
	response.SendResponse(c, resCode, gin.H{
		"id": cateDetail.Id,
		"name": cateDetail.Name,
		"desc": cateDetail.Desc,
		"create_time": cateDetail.CreateTime.Format("2006-01-02 15:04:05"),
	})
}
