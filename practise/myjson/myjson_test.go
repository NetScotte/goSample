package myjson

import (
	"fmt"
	"strings"
	"testing"
	"unicode"
)

func TestBasic(t *testing.T) {
	fmt.Println(Title("allow"))
	fmt.Println(Title("ALLOW"))
	fmt.Println(Title("ALL中文，OW"))
	fmt.Println(Title("中ALL中文，OW"))
	s := []rune("中Allow文")
	fmt.Printf("%T, %v\n", s[0], s[0])
}

func Title(s string) string {
	var r strings.Builder
	for index, i := range s {
		if index == 0 {
			r.WriteRune(unicode.ToUpper(i))
		} else {
			r.WriteRune(unicode.ToLower(i))
		}
	}
	return r.String()
}
