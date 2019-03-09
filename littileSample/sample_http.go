package main 

import (
	"fmt"
	"log"
	"net/http"
)

type Greeting string 

func (g Greeting) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, g)
}

func main() {
	err := http.ListenAndServe("localhost:8080", Greeting("Hello, http"))
	if err != nil {
		log.Fatal(err)
	}
}