package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// claim node : would declare properties that would be writed to neo4j
type neo4jNodeInfo struct {
	// testmap map[string]string
	DomainID string
	Name     string
	TAG      string
	Link     string
	Relation string
}

// var testString = "neo4j,nodeNum=n1,tag=testTag,domainID=testDomainID, name=testName,link=n1_n2|n3_n1,relation=belong|take,nodeNum=n2,tag=testTag2,domainID=testDomainID2, name=testName2,link=n2_n3,relation=follow,nodeNum=n3,tag=testTag3,domainID=testDomainID3, name=testName3,link=n1_n3,relation=after"

// assume data would only give neo4j input

// declare to parse json format while use NewRequest

func main() {
	// assume telegraf.Metric body would be like this
	// neo4j telegraf.Metric would be pass with map[name]field ; e.g. neo4j,nodeNum=n1,tag=testTag ,domainID=testDomainID, name=testName,link=n2_n3 (rmk: n2 belong to n3),relation=belong
	var neo4jInput = map[int]neo4jNodeInfo{
		1: {"testDomainID", "testName", "testTag1", "testDomainID_testDomainID2", "belong"},
		2: {"testDomainID2", "testName2", "testTag2", "testDomainID2_testDomainID3", "belong2"},
		3: {"testDomainID3", "testName3", "testTag3", "testDomainID3_testDomainID1", "belong3"},
	}

	// neo4jQueryNodes(neo4jInput["n1"])
	// neo4jQueryLink(neo4jInput["n1"])
	fmt.Println(neo4jQueryLink(neo4jInput[1]))

	// v,i:=range neo4jInput
	// fmt.Println(neo4jInput.len)

}

func neo4jQueryNodes(queryNodes neo4jNodeInfo) bool {
	type Payload struct {
		Query string `json:"query"`
	}
	data := Payload{
		Query: "match(n) where n.domainId='" + queryNodes.DomainID + "' return n.domainId",
	}
	// fmt.Println(queryNodes.DomainID)
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		// handle err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "http://127.0.0.1:7474/db/data/cypher", body)
	if err != nil {
		// handle err
	}
	req.SetBasicAuth("neo4j", "na")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	var f interface{}
	jsonResponse := f
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &jsonResponse)
	// print out jsonResponse
	fmt.Println(jsonResponse)

	criterion := (jsonResponse.(map[string]interface{})["data"].([]interface{})[0].([]interface{})[0] == queryNodes.DomainID)

	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
	return criterion
}

func neo4jQueryLink(queryNodes neo4jNodeInfo) bool {
	type Payload struct {
		Query string `json:"query"`
	}
	data := Payload{
		Query: "match p=(n1)-[r]->(n2) where type(r)='" + queryNodes.Relation + "' return n1.domainId,n2.domainId ,type(r)",
	}
	// fmt.Println(queryNodes.DomainID)
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		// handle err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "http://127.0.0.1:7474/db/data/cypher", body)
	if err != nil {
		// handle err
	}
	req.SetBasicAuth("neo4j", "na")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	var f interface{}
	jsonResponse := f
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &jsonResponse)

	criterion := (jsonResponse.(map[string]interface{})["data"].([]interface{})[0].([]interface{})[0] == queryNodes.DomainID)

	// value, criterion := jsonResponse.(map[string]interface{})["data"]
	// fmt.Println(value)

	// criterion := (jsonResponse.(map[string]interface{})["data"] != nil)

	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
	return criterion
	// return true
}
