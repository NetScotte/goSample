package main

import "fmt"

func changeParameter(array [3][]string) [3][]string {
	array[0] = []string{"changeParameter"}
	return array
}

func changeQuote(array [3][]string) [3][]string {
	array[0][0] = "changeQuote"
	return array
}

func main() {
	testarray := [3][]string{
		[]string{"a", "b", "c"},
		[]string{"d", "e", "f"},
		[]string{"1", "2", "3"},
	}
	fmt.Println("testarray is", testarray)
	fmt.Println("changeParameter--->", changeParameter(testarray))
	fmt.Println("testarray is", testarray)
	fmt.Println("changeQuote--->", changeQuote(testarray))
	fmt.Println("testarray is", testarray)
}
