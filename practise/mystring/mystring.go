package string

// 字符转置
func Reverse(s string) string {
	b := []rune(s)
	for i := 0; i < len(b)/2;i++ {
		j := len(b) - i - 1
		b[j], b[i] = b[i], b[j]
	}
	return string(b)
}