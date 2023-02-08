package myjson

// json示例
// Marshal(v any) ([]byte, error) 将数据结构转为数据,
// Unmarshal(data []byte, v any) error 将数据转为数据结构, 知道json时，定义对应的数据结构，为了通用，可以使用map[string]interface{}
// NewDecoder(r io.Reader) *Decoder   可以直接将io.Reader对象中的数据读到数据结构中, 适用于http中的body
// NewEncoder(w io.Writer) *Encoder   反之
// 将结构体转为json数据时，注意要指明json tag，否则可能是大写字母
// 如何解析key不固定的json内容？  map[string]object，string代表不固定的key，object代表固定的对象，
import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

// SampleUnmarshal 将数据格式转化为数据结构
func SampleUnmarshal() {
	var valid_json = []byte(`[
		{"name": "netliu", "age": 18}, 
		{"name": "freeze", "age": 22}
		]`)

	type Human struct {
		Name string
		Age  int
	}

	// 如果是`[{"name", "netliu", "age": 12}, {"name": "freeze", "age": 22}]`这种valid json
	var h []Human
	err := json.Unmarshal(valid_json, &h)
	if err != nil {
		fmt.Printf("%q", err)
	}
	fmt.Printf("%+v\n", h)

	// 如果是`{"name": "netliu", "age": 13}, {"name": "freeze", "age": 14}`这种invalid json
	var invalid_json = `
		{"name": "netliu", "age": 18}
		{"name": "freeze", "age": 22}
	`
	dec := json.NewDecoder(strings.NewReader(invalid_json))
	for {
		var sh Human
		if err := dec.Decode(&sh); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %d\n", sh.Name, sh.Age)
	}
}

// SampleMarshal 将数据结构转为数据
func SampleMarshal() {
	type Human struct {
		Name string
		Age  int
	}

	netliu := Human{
		Name: "netliu",
		Age:  12,
	}

	jsonStrings, err := json.Marshal(netliu)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("type: %T, value: %s\n", jsonStrings, jsonStrings)

	humans := []Human{
		{Name: "netliu", Age: 12},
		{Name: "freeze", Age: 22},
	}
	humanString, err := json.Marshal(humans)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("type: %T, value: %s\n", humanString, humanString)
}
