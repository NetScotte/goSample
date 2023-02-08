package myhttprequest

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func SampleHttp() {
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
	// 使用默认的http客户端可以复用TCP连接
	// http.DefaultClient.Do(req)
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

func BasicHTTPRequest(method, url string, requestData []byte) (responseData []byte, err error) {
	request, err := http.NewRequest(method, url, bytes.NewReader(requestData))
	if err != nil {
		return
	}
	client := &http.Client{
		Timeout: time.Second * 30,
	}
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("%s", p)
		}
	}()
	response, err := client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	log.Printf("status: %v\n", response.StatusCode)

	responseData, err = io.ReadAll(response.Body)
	if err != nil {
		return
	}
	log.Printf("%s\n", responseData)
	return
}
