package main

import (
	"log"
	"strings"
	"io/ioutil"
	"net/http"
)


type EasyApi struct{
	org string
	user	string
	cmdb_host string 

}

func (ea EasyApi)initHeader(header string) http.Header{
	Header := make(http.Header)
	Header.Add("Org", ea.org)
	Header.Add("User", ea.user)
	Header.Add("Content-type", "application/json")
	switch header {
	case "cmdb":
		Header.Add("Host", "cmdb.easyops-only.com")
	case "cmdb_resource":
		Header.Add("Host", "cmdb_resource.easyops-only.com")
	case "tool":
		Header.Add("Host", "cmdb.easyops-only.com")
	case "flow":
		Header.Add("Host", "easyflow.easyops-only.com")
	case "deploy":
		Header.Add("Host", "deploy.easyops-only.com")
	case "repo":
		Header.Add("Host", "deployrepo.easyops-only.com")
	case "dc_console":
		Header.Add("Host", "logic.dc_console.collector")
	}
	return Header
}


func (ea EasyApi)EasyRequest(url, method, header, params string) []byte {
	client := new(http.Client)
	body := strings.NewReader(params)
	req, _ := http.NewRequest(method, url, body)
	req.Header = ea.initHeader(header)
	req.Host = req.Header.Get("Host")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != 200 {
		log.Fatalf("StatusCode: %d, error: %s", resp.StatusCode, content)
	}
	return content
}

func (ea EasyApi) Get_all_instance(objectid, params string) string {
	url := "http://" + ea.cmdb_host + "/object/" + objectid + "/instance/_search"
	byte_result := ea.EasyRequest(url, "POST", "cmdb", params)
	string_result := string(byte_result)
	return string_result
}

func (ea EasyApi) Modify_instance(objectId, instanceId, params string) string {
	url := "http://" + ea.cmdb_host + "/object/" + objectId + "/instance/" + instanceId
	byte_result := ea.EasyRequest(url, "PUT", "cmdb", params)
	string_result := string(byte_result)
	return string_result
}


func main() {
	// query := `{"query": {"ip": "10.0.25.12"}, "fields": {"ip": 1, "hostname": 1}}`
	//cmdb_hosts := EasyApi("http://42.159.89.89/object/HOST/instance/_search", "POST", "cmdb", query)
	ea := EasyApi{
		org: "666003",
		user: "easyops",
		cmdb_host: "42.159.89.89",
	}
	// params := `{"query": {"ip": "10.0.25.12"}, "fields": {"ip": 1, "hostname": 1}}`
	// cmdb_hosts := ea.Get_all_instance("HOST", params)
	api_result := ea.Modify_instance("_IDCRACK", "5c6a4d30bee3f", `{"status": "test"}`)
	log.Printf("type: %T, value: %s", api_result, api_result)
}