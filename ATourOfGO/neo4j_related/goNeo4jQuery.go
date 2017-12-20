// refer to website
// https://stackoverflow.com/questions/38673673/access-http-response-as-string-in-go
// https://gobyexample.com/json

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func neo4jQueryNodes() AutoGenerated {

	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go
	type Payload struct {
		Query string `json:"query"`
	}

	data := Payload{
		// Query: "MATCH p=(n1)-[r]->(n2) return n1.domainId, type(r), n2.domainId",
		Query: "match(n) return n.domainId",
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		// handle err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "http://172.31.86.190:7474/db/data/cypher", body)
	if err != nil {
		// handle err
	}
	req.SetBasicAuth("neo4j", "na")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	res2 := AutoGenerated{}

	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)

		json.Unmarshal(bodyBytes, &res2)

	}

	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	return res2
}

type AutoGenerated struct {
	Columns []string        `json:"columns"`
	Data    [][]interface{} `json:"data"`
}

func main() {
	var test = neo4jQueryNodes()

	// Hereby the range test.Data doesn't mean it's contained value instead it's indexs
	if len(test.Data) >= 1 {
		for i := range test.Data {
				fmt.Println(test.Data[i])
		}
	}else{
		fmt.Println("length is 0")
	}
}