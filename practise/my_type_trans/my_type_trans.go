package my_type_trans

import (
	"strconv"
	"strings"
)

// 数字转字符
func IntToStr(a int) string {
	return strconv.Itoa(a)
}

// 字符转数字
func StrToInt(a string) (int, error) {
	s, err := strconv.Atoi(a)
	if err != nil {
		return 0, err
	}
	return s, nil
}

// 字符转分片
func StrToSlice(a string) []string {
	// []byte(a)
	return strings.Split(a, "")
}

// 分片转字符
func SliceToStr(a []string) string {
	return strings.Join(a, "")
}

// 字符转byte
func StrToByte(a string) []byte {
	return []byte(a)
}

// byte转字符
func ByteToStr(a []byte) string {
	return string(a)
}

func UintToInt(a uint) int {
	return int(a)
}

func IntToUint(a int) uint {
	return uint(a)
}

func ArrayToSlice(a [5]string) []string {
	return a[:]
}
