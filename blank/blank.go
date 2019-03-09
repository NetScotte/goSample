package main
import (
	"fmt"
	"encoding/json"
	"log"
)

func main() {
	params := make(map[string]string)
	query := `{ip: 10.0.25.12}`
	fields := `{instanceId: 1, hostname: 1}`
	params["page"] = "1"
	params["page_size"] = "300"
	fmt.Printf("query type: %T", query)
	if query != ""{
		params["query"] = query
	}
	if fields != "" {
		params["fields"] = fields  
	}

	string_params, err := json.Marshal(params)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("final params is: %s", string_params)
}