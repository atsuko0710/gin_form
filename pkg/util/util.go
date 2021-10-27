package util

// TruncateByWords 将帖子内容转换为简介
// s 帖子内容
// maxWords 转换后的字符长度
func TruncateByWords(s string, maxWords int) string {
	str := string([]rune(s)[:maxWords])
	res := str + "..."
	return res
}
