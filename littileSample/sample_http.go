package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	req, err := http.NewRequest("POST", "http://www.baidu.com/_search", nil)
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("haha", "haha")
	req.Header.Set("User", "user")
	req.Host = "www.baidu.com"
	// client := &http.Client{}
	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
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
