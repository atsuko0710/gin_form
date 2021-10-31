package enum

const (
	KeyPrefix          = "gin_forum:"
	KeyPostTime        = "post:time"   // zset;帖子以及发帖时间
	KeyPostScore       = "post:score"  // zset;帖子以及投票分数
	KeyPostVotedPrefix = "post:voted:" // zset;记录用户及投票类型
	KeyPostInfo        = "post:"       // hash;帖子详情，展示在首页
)
