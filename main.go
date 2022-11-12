package main

import (
	"fmt"
)

func EasyFunc(query interface{}, args ...interface{}) {
	fmt.Println("run")
	switch query.(type) {
	case []string:
		for _, arg := range args {
			switch arg := arg.(type) {
			case string:
				fmt.Printf("string %v\n", arg)
			case []string:
				fmt.Printf("[]string %v\n", arg)
			}
		}
	case string:
		fmt.Printf("string query %v\n", query)
	}
}

func main() {
	EasyFunc([]string{"aa", "bb"}, []string{"cc"})
}
