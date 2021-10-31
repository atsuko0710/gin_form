package util

// TruncateByWords 将帖子内容转换为简介
// s 帖子内容
// maxWords 转换后的字符长度
func TruncateByWords(s string, maxWords int) string {
	if len(s) > maxWords {
		str := string([]rune(s)[:maxWords])
		s = str + "..."	
	}
	return s
}
