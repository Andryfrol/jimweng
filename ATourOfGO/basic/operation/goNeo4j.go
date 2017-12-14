// This main code would use cypher to create nodes
// and deinfe function to create relation between nodes
// also both functions use http.Newrequest without .json file
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// gonna define the basic information of neo4j database
// including loginUser/ loginPasswd/ configURL
type writeInfo struct {
	loginUser   string
	loginPasswd string
	configURL   string
	command     string
}

// claim node : would declare properties that would be writed to neo4j
type nodeInfo struct {
	domainID  string
	name      string
	TAG       string
	link      string
	writeInfo writeInfo
}

func CreateNodes(c nodeInfo) {

	// Define create one node strings
	var oneNodeString = "create (n1:" + c.TAG + " {domainId:'" + c.domainID + "', name:'" + c.name + "'})"

	type Payload struct {
		Query string `json:"query"`
	}

	data := Payload{
		// fill struct
		Query: oneNodeString,
	}

	payloadBytes, err := json.Marshal(data)

	// fmt.Println(payloadBytes)

	if err != nil {
		// handle err
	}
	body := bytes.NewReader(payloadBytes)

	// fmt.Println(body)
	// fmt.Println("-------")
	// buf := new(bytes.Buffer) // create a temp memory to store
	// buf.ReadFrom(body)       // where r can be replace as any of Reader
	// s := buf.String()        // claim s as string for Reader
	// fmt.Println("Change body Reader to string :", s)

	req, err := http.NewRequest("POST", c.writeInfo.configURL, body)
	if err != nil {
		// handle err
	}

	req.SetBasicAuth(c.writeInfo.loginUser, c.writeInfo.loginPasswd)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
}

func ConnectNodes(c1, c2 nodeInfo) {
	// Define connect nodes strings
	var oneNodeString = "match (n1:" + c1.TAG + " {domainId:'" + c1.domainID + "', name:'" + c1.name + "'}) match (n2:" + c2.TAG + " {domainId:'" + c2.domainID + "', name:'" + c2.name + "'}) CREATE p2 = (n2)-[:BELONG]->(n1) RETURN p2"

	type Payload struct {
		Query string `json:"query"`
	}

	data := Payload{
		// fill struct
		Query: oneNodeString,
	}

	payloadBytes, err := json.Marshal(data)

	// fmt.Println(payloadBytes)

	if err != nil {
		// handle err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", c1.writeInfo.configURL, body)
	if err != nil {
		// handle err
	}

	req.SetBasicAuth(c1.writeInfo.loginUser, c1.writeInfo.loginPasswd)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
}

func CreateConnectedNodes(c1, c2 nodeInfo) {
	// Define connect nodes strings
	var oneNodeString = "create (n1:" + c1.TAG + " {domainId:'" + c1.domainID + "', name:'" + c1.name + "'}) create (n2:" + c2.TAG + " {domainId:'" + c2.domainID + "', name:'" + c2.name + "'}) CREATE p2 = (n2)-[:BELONG]->(n1) RETURN p2"

	type Payload struct {
		Query string `json:"query"`
	}

	data := Payload{
		// fill struct
		Query: oneNodeString,
	}

	payloadBytes, err := json.Marshal(data)

	// fmt.Println(payloadBytes)

	if err != nil {
		// handle err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", c1.writeInfo.configURL, body)
	if err != nil {
		// handle err
	}

	req.SetBasicAuth(c1.writeInfo.loginUser, c1.writeInfo.loginPasswd)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
}

func CreateNodes(c []nodeInfo) {

	// Define create one node strings
	var oneNodeString = "create (n1:" + c.TAG + " {domainId:'" + c.domainID + "', name:'" + c.name + "'})"

	type Payload struct {
		Query string `json:"query"`
	}

	data := Payload{
		// fill struct
		Query: oneNodeString,
	}

	payloadBytes, err := json.Marshal(data)

	// fmt.Println(payloadBytes)

	if err != nil {
		// handle err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", c.writeInfo.configURL, body)
	if err != nil {
		// handle err
	}

	req.SetBasicAuth(c.writeInfo.loginUser, c.writeInfo.loginPasswd)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
}

func main() {

	commit1 := nodeInfo{
		TAG:      "VMware",
		domainID: "172.31.1.1",
		name:     "jim",
		writeInfo: writeInfo{
			loginUser:   "neo4j",
			loginPasswd: "na",
			configURL:   "http://172.31.86.190:7474/db/data/cypher",
		},
	}

	commit2 := nodeInfo{
		TAG:      "VMware",
		domainID: "172.31.0.1",
		name:     "bob",
		writeInfo: writeInfo{
			loginUser:   "neo4j",
			loginPasswd: "na",
			configURL:   "http://172.31.86.190:7474/db/data/cypher",
		},
	}

	if commit1.link == "" {
		fmt.Println("no link")
	}
	// create one node
	// CreateNodes(commit1)
	// CreateNodes(commit2)

	// create connect between nodes
	// ConnectNodes(commit1, commit2)

	// create nodes and connection at the same time
	// CreateConnectedNodes(commit1, commit2)

	// create mulitple nodes

}
