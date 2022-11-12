package myset

import (
	"fmt"
	"testing"
)

func TestStringSet_Add(t *testing.T) {
	set := NewStringSet()
	set.Add("a")
	set.Add("a")
	set.Add("3")
	set.Add("c")
	fmt.Println(set)
}

func TestSet_String(t *testing.T) {
	set := NewSet()
	set.Add("a")
	set.Add("a")
	set.Add("3")
	set.Add("c")
	fmt.Println(set)
}

func TestSet_Int(t *testing.T) {
	set := NewSet()
	set.Add(1)
	set.Add(2)
	set.Add(2)
	set.Add(4)
	fmt.Println(set)
}

func TestSet_Float(t *testing.T) {
	set := NewSet()
	set.Add(1.2)
	set.Add(2.2)
	set.Add(2.2)
	set.Add(4.1)
	fmt.Println(set)
}
