package service

import (
	"database/sql"
	"fmt"
	"gin_forum/config/redis"
	"gin_forum/models"
	"gin_forum/params"
	"gin_forum/pkg/enum"
	"gin_forum/pkg/response"
	"gin_forum/repository"
	redisDriver "github.com/go-redis/redis"
	"strconv"
	"time"

	"go.uber.org/zap"
)

const POST_NORMAL_STATUS = 1

func CreatePost(request params.CreatePostRequest) response.ResCode {
	categoryId, err := strconv.ParseInt(request.CategoryId, 10, 64)
	if err != nil {
		return response.InvalidParam
	}

	post := models.Post{
		Title:      request.Title,
		Content:    request.Content,
		CategoryId: categoryId,
		AuthorId:   request.AuthorId,
		Status:     POST_NORMAL_STATUS,
	}

	res, err := repository.CreatePost(post)
	if err != nil {
		return response.CreatePostFail
	}

	err = createPostRedis(res)
	if err != nil {
		zap.L().Error("post.createPostRedis() fail", zap.Error(err))
		return response.InvalidParam
	}
	return response.OK
}

/*
createPostRedis 创建投票缓存

这里采用 《redis 实战》中投票的例子
规定一篇文章获得 200 张支持票后就能认为是一篇有趣的文章，应该将这篇文章放到首页至少一天
则 86400 / 200 = 432 粗略计算一票为 432分
为了产生一个随着时间流逝分数不断变少的的评分效果，则直接将新发布的文章设置初始分数为当前时间戳
*/
func createPostRedis(post models.Post) (err error) {
	now := float64(time.Now().Unix())
	voteKey := enum.KeyPostVotedPrefix + fmt.Sprint(post.Id)

	// postInfo := map[string]interface{}{
	// 	"title":    post.Title,
	// 	"summary":  util.TruncateByWords(post.Content, 120),
	// 	"post:id":  post.Id,
	// 	"user:id":  post.AuthorId,
	// 	"time":     now,
	// 	"votes":    1,
	// 	"comments": 0,
	// }

	// 事务操作
	pipeline := redis.Client.TxPipeline()
	pipeline.ZAdd(voteKey, redisDriver.Z{ // 发布帖子的作者默认投一票
		Score:  1,
		Member: post.AuthorId,
	})

	pipeline.ZAdd(enum.KeyPostScore, redisDriver.Z{ // 增加默认一票的分数
		Score:  now + VoteScore,
		Member: post.Id,
	})

	pipeline.ZAdd(enum.KeyPostTime, redisDriver.Z{  // 添加时间
		Score:  now,
		Member: post.Id,
	})
	_, err = pipeline.Exec()
	return
}

func GetPostDetail(Id int64) (c *params.PostDetailResponse, resCode response.ResCode) {
	c, err := repository.GetPostDetail(Id)

	// 没有找到结果
	if err == sql.ErrNoRows {
		return &params.PostDetailResponse{}, response.GetDetailFail
	}

	if err != nil {
		zap.L().Error("repository.GetPostDetail() failed", zap.Error(err))
		return &params.PostDetailResponse{}, response.GetDetailFail
	}
	return c, response.OK
}

func GetPostList(index int64, count int64) (posts []params.ApiPostDetailResponse) {
	posts = repository.GetPostList(index, count)
	return
}
