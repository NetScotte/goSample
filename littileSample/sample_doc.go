package main 

import (
	"fmt"
)

func ExampleIndex() {
	fmt.Println(strings.Index("chiken", "ken"))
	fmt.Println(strings.Index("chiken", "dmr"))
	// Output:
	// 4
	// -1
}