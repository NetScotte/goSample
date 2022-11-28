package myhttprequest

import (
	"encoding/json"
	"testing"
)

func TestHttpRequestSample(t *testing.T) {
	SampleHttp()
}

func TestBasicHTTPRequest(t *testing.T) {
	data, _ := json.Marshal(map[string]string{
		"status": "on",
	})
	content, err := BasicHTTPRequest("GET", "http://localhost:8181/get?name=http", data)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(string(content))
	}
}
