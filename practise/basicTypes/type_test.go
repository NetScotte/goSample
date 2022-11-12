package basicTypes

import (
	"fmt"
	"testing"
)

// Printf提供：
// %s 字符串
// %T   类型
func TestPoint(t *testing.T) {
	fmt.Println("pointer_sample")
	// define a pointer
	var p *int
	// define a int type
	number := 4
	fmt.Println("number:", number, "&number:", &number)
	// give value to pointer
	p = &number
	// show pointer content
	fmt.Println("pointer p:", p, "*p:", *p)
	// change point content
	*p = *p + 1
	fmt.Println("p:", p, "*p", *p)
	fmt.Println()
}

func TestNil(t *testing.T) {
	fmt.Println("nil_sample:")
	// use nil direct
	fmt.Println("nil: ", nil)
	// define "", if not use, error occur
	// blank := ""
	// mismatch
	// fmt.Println("nil==blank: ", nil==blank)
	fmt.Println()
}

func TestStruct(t *testing.T) {
	fmt.Println("strcut_sample:")
	// define struct
	type People struct {
		name string
		age  int
		man  bool
	}
	// instance struct
	p := People{"lfy", 25, true}
	// show struct
	fmt.Println("p: ", p)
	// change struct content
	p.age += 1
	fmt.Println("p.age += 1, p: ", p)
	fmt.Println()
}

func TestArray(t *testing.T) {
	// define a array
	var array [2]string
	var array2 = [2]string{"a", "b"}
	var array3 = make([]int, 3)
	fmt.Printf("array2: %T, %v\n", array2, array2)
	fmt.Printf("array3: %T, %v\n", array3, array3)
	// use array
	array[0] = "Hello"
	array[1] = "World"
	// array is [Hello World]
	fmt.Println("array is: ", array)
	// change content
	array[1] = "lfy"
	fmt.Println("array is: ", array)
	fmt.Println("array len: ", len(array))
	fmt.Println("array cap: ", cap(array))
	fmt.Println()
}

func TestSlice(t *testing.T) {
	s := []int{1, 2, 3, 4}
	s[0] = 0
	m := make([]int, 0, 5)
	a := append(m, 1, 2, 3)
	fmt.Println("m is: ", m)
	fmt.Println("a is: ", a)
	// error operator
	// p := s + m
	// for _, value := range s {
	//     a = append(a, value)
	// }
	// 优雅得合并
	a = append(a, s...)
	fmt.Println("s is: ", s)
	fmt.Println("a is: ", a)
}

func TestMap(t *testing.T) {
	// define map, error sample: m = make(map[string]string ; m := map[string]string
	m := make(map[string]string)

	// set value for map
	m["wh"] = "wuhan"
	m["bj"] = "beijing"
	fmt.Println("m is: ", m)

	// change map
	m["wh"] = "WuHan"
	fmt.Println("m is: ", m)

	// check map key, 'wh' is rune
	value, ok := m["wh"]

	// can't use %s for ok
	fmt.Println("value: ", value, "ok: ", ok)

	// 获取长度
	fmt.Println(len(m))

	// delete map key
	delete(m, "bj")

	// 获取长度
	fmt.Println(len(m))

}
