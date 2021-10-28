package service

import (
	"gin_forum/config/redis"
	"gin_forum/params"
	"gin_forum/pkg/enum"
	"gin_forum/pkg/response"
	"math"
	"time"

	redisDriver "github.com/go-redis/redis"
	"github.com/spf13/cast"
)

const (
	OneWeekInSeconds = 7 * 24 * 3600 // 一周秒数
	VoteScore        = 432           // 一票获取分数
)

/*
投票分为四种情况：1.投赞成票 2.投反对票 3.取消投票 4.反转投票

记录文章参与投票的人
更新文章分数：赞成票要加分；反对票减分

v=1时，有两种情况
	1.之前没投过票，现在要投赞成票  => +432
	2.之前投过反对票，现在要改为赞成票  => +432*2
v=0时，有两种情况
	1.之前投过赞成票，现在要取消 => -432
	2.之前投过反对票，现在要取消  => +432
v=-1时，有两种情况
	1.之前没投过票，现在要投反对票 => -432
	2.之前投过赞成票，现在要改为反对票 => -432*2
*/
func VoteForPost(userId string, v params.VoteRequest) (resCode response.ResCode) {
	// 获取帖子发布时间
	postTime := redis.Client.ZScore(enum.KeyPostTime, v.PostId).Val()

	// 当帖子超过一周后就不再投票了
	if float64(time.Now().Unix())-postTime > OneWeekInSeconds {
		return response.VoteTimeExpire
	}

	voteKey := enum.KeyPostVotedPrefix + v.PostId
	ov := redis.Client.ZScore(voteKey, userId).Val() // 获取当前用户投票记录
	diffAbs := math.Abs(ov - v.Vote)

	pipeline := redis.Client.TxPipeline()

	pipeline.ZAdd(voteKey, redisDriver.Z{ // 更新投票记录
		Score:  v.Vote,
		Member: userId,
	})

	pipeline.ZIncrBy(enum.KeyPostScore, VoteScore*diffAbs*v.Vote, v.PostId) // 更新分数

	switch math.Abs(ov) - math.Abs(v.Vote) {
	case 1:
		pipeline.HIncrBy(K)
	}

	return response.OK
}
