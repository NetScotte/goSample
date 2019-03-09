package main 

import (
	"fmt"
	"time"
)

func main() {
	// sample 1: current time
	t := time.Now()
	fmt.Prinln(t.In(time.UTC))
	home, _ = time.LoadLocation("ShangHai/Asia")
	fmt.Prinln(t.In(home))

	if time.Now().Hour() < 12 {
		fmt.Println("Good morning")
	} else {
		fmt.Println("Good afternoon (or evening)")
	}

	// sample 2: time calculate
	birthday, _ = time.Parse("Oct 31 1993", "Jan 19 2019")
	age, _ = time.Since(birthday)
	fmt.Printf("Net is %d days old", age/(time.Hour*24))
}