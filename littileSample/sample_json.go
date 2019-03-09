package main 

import (
	"fmt"
	"encoding/json"
	"strings"
)

const blob = '[
	{"Title": "redev", "URL": "http://oredev.org"},
	{"Title": "Strange Loop", "URL": "http://thestrangeloop.com"}
]'

type Item struct {
	Title string 
	URL string 
}

func main() {
	var items []*Item 
	json.NewDecode(strings.NewReader(blob)).Decode(&items)
	for _, item in range items {
		fmt.Println("URL: %v \n", item.URL)
	}
}