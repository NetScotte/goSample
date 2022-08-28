package myreflect

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string `color:"Red" json:"name"`
	Age   int
	Class string
}

func (s *Student) GetName() string {
	return s.Name
}

func (s *Student) GetAge() int {
	return s.Age
}

func (s *Student) SetName(name string) {
	s.Name = name
}

func BasicTypeWithReflect() {
	i := 3
	f := 3.8
	// 获取类型和值
	fmt.Printf("i type: %v, value: %v\n", reflect.TypeOf(i), reflect.ValueOf(i))
	fmt.Printf("f type: %v, value: %v\n", reflect.TypeOf(f), reflect.ValueOf(f))
	fmt.Println("set value for i = 4, f = 6.8")
	reflect_value_of_i := reflect.ValueOf(&i)
	// 修改对象的值
	reflect_value_of_i.Elem().SetInt(4)
	reflect_valueof_f := reflect.ValueOf(&f)
	reflect_valueof_f.Elem().SetFloat(6.8)
	fmt.Printf("now i=%d, f=%f\n", i, f)
}

func StructTypeWithReflect() {
	fmt.Println("")
	stu := Student{"netliu", 25, "College"}
	fmt.Printf("stu type: %v, value: %v\n", reflect.TypeOf(stu), reflect.ValueOf(stu))
	fmt.Println("get all attribute of struct")
	// type类型没有方法, 获取结构体属性信息
	reflect_value_of_stu := reflect.ValueOf(stu)
	reflect_type_of_stu := reflect.TypeOf(stu)
	for i := 0; i < reflect_value_of_stu.NumField(); i++ {
		value := reflect_value_of_stu.Field(i)
		t := reflect_type_of_stu.Field(i)
		fmt.Printf("attribute name: %v, type: %v, value: %v\n", t.Name, value.Type(), value.Interface())
	}
	// 设置结构体的值
	fmt.Println("set stu age=20")
	reflect_ptr_value_of_stu := reflect.ValueOf(&stu)
	reflect_ptr_type_of_stu := reflect.TypeOf(&stu)
	reflect_ptr_value_of_stu.Elem().Field(1).SetInt(20)
	fmt.Printf("stu is %v\n", stu)

	// 获取接替的函数信息
	fmt.Printf("[error]stu has %v method\n", reflect_value_of_stu.NumMethod())
	fmt.Printf("stu has %v method\n", reflect_ptr_value_of_stu.NumMethod())
	for i := 0; i < reflect_ptr_value_of_stu.NumMethod(); i++ {
		method_field_type := reflect_ptr_type_of_stu.Method(i)
		fmt.Printf("method name: %v, type: %v\n", method_field_type.Name, method_field_type.Type)
	}

	// 调用结构体的方法
	fmt.Println("set name to Fang")
	stu_setname_method := reflect_ptr_value_of_stu.MethodByName("SetName")
	var params []reflect.Value
	name := "Fang"
	params = append(params, reflect.ValueOf(name))
	stu_setname_method.Call(params)
	fmt.Printf("stu is %v\n", stu)

	// 获取结构体的tag, 不能通过指针类型获取值属性
	for i := 0; i < reflect_value_of_stu.NumField(); i++ {
		type_field := reflect_type_of_stu.Field(i)
		fmt.Printf("attribute name: %v with tag: %v, color tag value: %v\n", type_field.Name, type_field.Tag, type_field.Tag.Get("color"))
	}

}

type NewStudent struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func interfaceSample(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("int")
	default:
		fmt.Println("")
	}
}

func sampleTotal() {
	var a int = 3
	// 空接口
	fmt.Println("interface sample -->")
	interfaceSample(a)

	// reflect 基础
	fmt.Println("\nreflect basic sample -->")
	aType := reflect.TypeOf(a)   // 获取类型
	aValue := reflect.ValueOf(a) // 获取值
	fmt.Println(aType)           // int
	fmt.Println(aValue)          // 3

	// 转为真实对象
	if aType.Kind() == reflect.Int {
		aReal := aValue.Int()
		fmt.Println("= 3 ? ", aReal == 3)
	}

	// 操作数组或分片
	fmt.Println("\narray or slice sample -->")
	var b = []int{1, 2, 3, 4, 5}
	bValue := reflect.ValueOf(b)
	bType := reflect.TypeOf(b)
	fmt.Println(bType, bType.Kind())
	for i := 0; i < bValue.Len(); i++ { // 遍历数组
		fmt.Print(bValue.Index(i), " ")
	}
	fmt.Println()

	// 操作结构体
	fmt.Println("\nstruct sample -->")
	var s = NewStudent{"net", 10}
	sValue := reflect.ValueOf(s)
	sType := reflect.TypeOf(s)
	fmt.Println(sType, sType.Kind())
	for i := 0; i < sValue.NumField(); i++ {
		item := sValue.Field(i)
		// 无法通过item获取Tag, 必须先获取类型
		fmt.Println(item.Type(), item, sType.Field(i).Tag.Get("json"))
	}

}
