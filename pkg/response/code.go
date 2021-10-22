package response

type ResCode int64

const (
	OK ResCode = 1000 + iota
	InvalidParam
	UserExist
	UserNotExist
	InvalidPassword
	CreateUserField
	InternalServerError
	InvalidToken
)

var codeMsgMap = map[ResCode]string{
	OK:                  "success",
	InvalidParam:        "请求参数错误",
	UserExist:           "用户已存在",
	UserNotExist:        "用户名不存在",
	InvalidPassword:     "用户名或密码错误",
	CreateUserField:     "创建用户失败",
	InternalServerError: "服务器错误",
	InvalidToken:        "生成Token失败",
}

// Msg 返回错误描述
func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[InternalServerError]
	}
	return msg
}
