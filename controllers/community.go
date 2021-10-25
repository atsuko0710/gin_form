package controllers

import (
	"gin_forum/params"
	"gin_forum/pkg/response"
	"gin_forum/service"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// CategoryList 分类列表入口
func CategoryList(c *gin.Context) {
	cateList, resCode := service.GetCategoryList()
	response.SendResponse(c, resCode, cateList)
}

// CategoryDetail 分类详情入口
func CategoryDetail(c *gin.Context) {
	idStr := c.Param("id")
	CategoryId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.SendResponse(c, response.InvalidParam, gin.H{})
	}
	cateDetail, resCode := service.GetCategoryDetail(CategoryId)
	response.SendResponse(c, resCode, gin.H{
		"id":          cateDetail.Id,
		"name":        cateDetail.Name,
		"desc":        cateDetail.Desc,
		"create_time": cateDetail.CreateTime.Format("2006-01-02 15:04:05"),
	})
}

// CreatePost 创建帖子入口
func CreatePost(c *gin.Context) {
	var params params.CreatePostRequest
	authorId, _, _ := getCurrentUser(*c)
	if err := c.ShouldBindJSON(&params); err != nil {
		// 判断错误是不是 validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			zap.L().Error("CreatePost.validator() failed", zap.Error(err))
			response.SendResponse(c, response.InvalidParam, gin.H{})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}
	params.AuthorId = authorId
	resCode := service.CreatePost(params)
	response.SendResponse(c, resCode, gin.H{})
}

// PostDetail 获取帖子详情
func PostDetail(c *gin.Context) {
	idStr := c.Param("id")
	PostId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.SendResponse(c, response.InvalidParam, gin.H{})
	}
	postDetail, resCode := service.GetPostDetail(PostId)
	response.SendResponse(c, resCode, gin.H{
		"id":          postDetail.Id,
		"title":       postDetail.Title,
		"content":     postDetail.Content,
		"author_id":   postDetail.AuthorId,
		"status":      postDetail.Status,
		"category_id": postDetail.CategoryId,
		"create_time": postDetail.CreateTime.Format("2006-01-02 15:04:05"),
	})
}
