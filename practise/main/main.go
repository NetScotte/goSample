package main

import (
	_ "errors"
	"fmt"

	_ "github.com/pkg/errors"
)

func main() {
	err := test0()
	fmt.Printf("%+v", err)
}

func test0() error {
	return test1()
}

func test1() error {
	return test2()
}

func test2() error {
	return fmt.Errorf("%v", "a error in test2")
}
