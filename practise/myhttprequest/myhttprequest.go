package myhttprequest

import (
	"io/ioutil"
	"log"
	"net/http"
)

func HttpRequestSample() {
	// 创建请求
	req, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	if err != nil {
		log.Fatal(err)
	}
	// 创建客户端
	client := new(http.Client)
	// 发起请求
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	// 解析结果
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("failed read content from res.Body")
	}
	log.Printf("status: %v\n", res.StatusCode)
	log.Printf("%s\n", content)

}
