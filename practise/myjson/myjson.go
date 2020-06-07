package myjson

import (
	"fmt"
	"encoding/json"
	"strings"
	"io"
	"log"
)


func Sample_json2struct() {
	var valid_json = []byte(`[
		{"name": "netliu", "age": 18}, 
		{"name": "freeze", "age": 22}
		]`)

	type Human struct {
		Name string
		Age int 
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
		}else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %d\n", sh.Name, sh.Age)
	}
}

func Sample_struct2json() {
	type Human struct {
		Name string 
		Age		int 
	}
	fmt.Println("Sample_struct2json.............Human to string")
	netliu := Human{
		Name: "netliu",
		Age:	12,
	}
	// 无法将humans通过string进行强转
	jsonstrings, err := json.Marshal(netliu)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("type: %T, value: %s\n", jsonstrings, jsonstrings)
	fmt.Println("Sample_struct2json...........[]Human to string")

	humans := []Human{
		{Name: "netliu", Age: 12},
		{Name: "freeze", Age: 22},
	}
	humansting, err := json.Marshal(humans)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("type: %T, value: %s\n", humansting, humansting)
}


func Sample_json2sturct2() {
	rawJson := `{"name": "netliu", "age": 12, "sex": "man"}`
	type Man struct{

	}
	m := new(Man)
	err := json.Unmarshal([]byte(rawJson), m)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("m is: %+v\n", *m)
}