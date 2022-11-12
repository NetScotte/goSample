package myset

import (
	"fmt"
	"testing"
)

func TestStringSet_Add(t *testing.T) {
	set := NewStringSet()
	set.Add("1")
	set.Add("a")
	set.Add("3")
	set.Add("c")
	fmt.Println(set)
}
