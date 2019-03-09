package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
)

func main() {
	req, err := http.NewRequest("POST", "http://42.159.89.89/object/HOST/instance/_search", nil)
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Org", "666003")
	req.Header.Set("User", "easyops")
	req.Host = "cmdb.easyops-only.com"
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		fmt.Println(req.URL)
		fmt.Println(req.Header)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("url: %s\n", req.URL)
	fmt.Printf("header: %s\n", req.Header)
	fmt.Printf("result: %s\n", body)
	
}