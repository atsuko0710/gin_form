package repository

import (
	"context"
	"gin_forum/config/elasticsearch"
	"gin_forum/models"
	"gin_forum/params"
	"reflect"

	"github.com/olivere/elastic/v7"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func CreatePost(p models.Post) (post models.Post, err error) {
	return models.CreatePost(p)
}

func GetPostDetail(Id int64) (c *params.PostDetailResponse, err error) {
	c, err = models.FindPost(Id)
	return
}

func GetPostList(index int64, count int64) (posts []params.ApiPostDetailResponse) {
	posts = models.FindPostWithRelate(index, count)
	return
}

// CreatePostInES 将 post 数据更新到 es 用来查询
func CreatePostInES(p params.PostSearchResponse) (err error) {
	res, err := elasticsearch.Client.Index().
		Index(viper.GetString("name")).
		Type("post").
		Id(p.PostId).
		BodyJson(p).
		Do(context.Background())
	zap.L().Info("Indexed es result", zap.String("Id", res.Id), zap.String("Index", res.Index), zap.String("Type", res.Type))

	if err != nil {
		zap.L().Error("elasticsearch.Index() Fail", zap.Error(err))
	}
	return
}

// SearchPostInES 在 ES 查询 post 数据
func SearchPostInES(param string) []params.PostSearchResponse {
	boolQuery := elastic.NewBoolQuery().Must()

	Param1 := elastic.NewMatchPhraseQuery("title", param)
	Param2 := elastic.NewMatchPhraseQuery("content", param)
	boolQuery.Should(Param1, Param2)

	matchResult, err := elasticsearch.Client.
		Search().
		Index(viper.GetString("name")).
		Type("post").
		Query(boolQuery).
		Do(context.Background())

	if err != nil {
		zap.L().Error("elasticsearch.Search() Fail", zap.Error(err))
	}
	res := getResultInEs(matchResult)
	return res
}

// getResultInEs 将 es 结果按规范返回
func getResultInEs(res *elastic.SearchResult) []params.PostSearchResponse {
	var postSearch params.PostSearchResponse
	var result []params.PostSearchResponse
	for _, item := range res.Each(reflect.TypeOf(postSearch)) {
		p := item.(params.PostSearchResponse)
		zap.L().Info("getResultInEs", zap.String("title", p.Title))
		result = append(result, p)
	}
	return result
}
