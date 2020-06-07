package mylock

import (
	"sync"
	"testing"
)

func TestAll(t *testing.T) {
	var wg sync.WaitGroup
	u := User{name: "netliu"}
	wg.Add(1)
	go Deposit(&u)
	go Draw(&u)
	go monitor(&u, &wg)
	wg.Wait()
}
