package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

func curlWithJSONFile() {
	f, err := os.Open("create.json")
	if err != nil {
		// handle err
	}
	defer f.Close()
	req, err := http.NewRequest("POST", "http://172.31.86.178:7474/db/data/cypher", f)
	if err != nil {
		// handle err
	}
	req.SetBasicAuth("neo4j", "password")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
}

func curlwithoutJSONfile() {
	// claim node
	var nameString = "john"
	var queryString = "create (n:Person {name:'" + nameString + "'})"

	type Payload struct {
		Query string `json:"query"`
	}

	data := Payload{
		// fill struct
		Query: queryString,
	}

	payloadBytes, err := json.Marshal(data)

	if err != nil {
		// handle err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "http://172.31.86.178:7474/db/data/cypher", body)
	if err != nil {
		// handle err
	}
	req.SetBasicAuth("neo4j", "password")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
}

func main() {
	curlwithoutJSONfile()
	// curlWithJSONFile()
}
