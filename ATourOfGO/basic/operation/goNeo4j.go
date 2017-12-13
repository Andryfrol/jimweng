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
	nodeInfor   nodeInfor
}

// claim node : would declare properties that would be writed to neo4j
type nodeInfor struct {
	domainID string
	name     string
	TAG      string
	connect  bool
}

// // neo4j interface
// // would be designed to isolate different data
// type neo4jDB interface {
// 	CreateNodes()
// 	QueryNodes()
// 	ConnectNodes()
// }

func CreateNodes(c writeInfo) {

	// Define create one node strings
	var oneNodeString = "create (n1:" + c.nodeInfor.TAG + " {domainId:'" + c.nodeInfor.domainID + "', name:'" + c.nodeInfor.name + "'})"

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

	req, err := http.NewRequest("POST", c.configURL, body)
	if err != nil {
		// handle err
	}

	req.SetBasicAuth(c.loginUser, c.loginPasswd)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
}

func main() {

	commit1 := writeInfo{
		loginUser:   "neo4j",
		loginPasswd: "na",
		configURL:   "http://172.31.86.190:7474/db/data/cypher",
		nodeInfor: nodeInfor{
			TAG:      "VMware",
			domainID: "172.31.1.1",
			name:     "Jim",
			connect:  false,
		},
	}
	fmt.Println(commit1.nodeInfor)

	CreateNodes(commit1)
}
