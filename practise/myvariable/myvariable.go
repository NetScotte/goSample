package main
import "fmt"


// Printf提供：
// %s 字符串
// %T   类型
func pointer_sample() {
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

func nil_sample() {
	fmt.Println("nil_sample:")
	// use nil direct
	fmt.Println("nil: ", nil)
	// define "", if not use, error occur
	// blank := ""
	// mismatch
	// fmt.Println("nil==blank: ", nil==blank)
	fmt.Println()
}

func struct_sample() {
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

func array_sample() {
	// define a array
	var array [2]string
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

func slice_sample() {
	s := []int{1,2,3,4}
	s[0]=0
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

func map_sample() {
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

	// delete map key
	delete(m, "bj")
}

func main() {
	pointer_sample()
	nil_sample()
	struct_sample()
	array_sample()
	slice_sample()
	map_sample()
}
