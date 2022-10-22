package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	rawJson := `{"name": "netliu", "age": 12, "sex": "man"}`
	type Man struct {
	}
	m := new(Man)
	err := json.Unmarshal([]byte(rawJson), m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("m is: %+v\n", *m)

}
