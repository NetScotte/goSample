package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	for index, value := range a {
		fmt.Printf("index: %d, value: %d\n", index, value)
	}
}
