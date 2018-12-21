package main
import "fmt"

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
    
func main() {
    pointer_sample()
    nil_sample()
    struct_sample()
    array_sample()
}
