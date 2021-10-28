package response

type ResCode int64

const (
	OK ResCode = 1000 + iota
	InvalidParam
	UserExist
	UserNotExist
	InvalidPassword
	CreateUserFail
	InternalServerError
	InvalidToken
	UserNotLogin
	GetListFail
	GetDetailFail
	CreatePostFail
	VoteTimeExpire
)

var codeMsgMap = map[ResCode]string{
	OK:                  "success",
	InvalidParam:        "请求参数错误",
	UserExist:           "用户已存在",
	UserNotExist:        "用户名不存在",
	InvalidPassword:     "用户名或密码错误",
	CreateUserFail:      "创建用户失败",
	InternalServerError: "服务器错误",
	InvalidToken:        "生成Token失败",
	UserNotLogin:        "用户未登陆",
	GetListFail:         "获取列表信息失败",
	GetDetailFail:       "获取详情信息失败",
	CreatePostFail:      "创建帖子失败",
	VoteTimeExpire:      "投票时间已超时",
}

// Msg 返回错误描述
func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[InternalServerError]
	}
	return msg
}
