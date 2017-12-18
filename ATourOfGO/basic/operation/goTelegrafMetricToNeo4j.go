package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
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

func CreateMultiNodes(cA []nodeInfo) {
	// var actString = ""
	for index := 0; index < len(cA); index++ {
		// fmt.Println(index)
		c := cA[index]
		// if c.writeInfo.command == "CreateNodes" {
		// actString = "create (n1:" + c.TAG + " {domainId:'" + c.domainID + "', name:'" + c.name + "'})"
		CreateNodes(c)
		// }
		// fmt.Println(c.writeInfo.command) //check command
	}
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

	// assume telegraf.Metric body would be like this
	var testString = "mem,host=MacPro used_percent=57.654523849487305,total=8589934592i,cached=0i,inactive=3101143040i,slab=0i,active=3521380352i,available_percent=42.345476150512695,available=3637448704i,used=4952485888i,free=536305664i,buffered=0i 1513149940000000000\ndisk,path=/,device=disk1s1,fstype=apfs,mode=rw,host=MacPro inodes_free=9223372036853631618i,inodes_used=1144189i,total=119824367616i,free=71182340096i,used=42605252608i,used_percent=37.44279283491888,inodes_total=9223372036854775807i 1513149940000000000\nsystem,host=MacPro load15=3.35,n_users=1i,n_cpus=4i,load1=4.48,load5=4.86 1513566990000000000\nsystem,host=MacPro uptime=6809i,uptime_format= 1513566990000000000"

	// define commitArray
	// var commitArray1 = []nodeInfo{}
	// split body into measuremets
	var mesurementStringArray = strings.Split(testString, "\n")
	// create request body also make criterion to choose

	for index := 0; index < len(mesurementStringArray); index++ {
		neededMeasurement := strings.Split(mesurementStringArray[index], ",")
		if neededMeasurement[0] == "system" || neededMeasurement[0] == "mem" {

			nodesCreated := nodeInfo{
				TAG:      "VMware",
				domainID: neededMeasurement[1],
				name:     neededMeasurement[0],
				writeInfo: writeInfo{
					loginUser:   "neo4j",
					loginPasswd: "na",
					configURL:   "http://172.31.86.190:7474/db/data/cypher",
					command:     "CreateNodes",
				},
			}
			// fmt.Println(mesurementStringArray)
			CreateNodes(nodesCreated)
		}
	}

	// var commitArray2 = []nodeInfo{
	// 	nodeInfo{
	// 		TAG:      "VM",
	// 		domainID: "172.31.1.1",
	// 		name:     "amy",
	// 		writeInfo: writeInfo{
	// 			loginUser:   "neo4j",
	// 			loginPasswd: "na",
	// 			configURL:   "http://172.31.86.190:7474/db/data/cypher",
	// 			command:     "CreateNodes",
	// 		},
	// 	},

	// 	nodeInfo{
	// 		TAG:      "VM",
	// 		domainID: "172.31.0.1",
	// 		name:     "kira",
	// 		writeInfo: writeInfo{
	// 			loginUser:   "neo4j",
	// 			loginPasswd: "na",
	// 			configURL:   "http://172.31.86.190:7474/db/data/cypher",
	// 			command:     "CreateNodes",
	// 		},
	// 	},
	// 	nodeInfo{
	// 		TAG:      "VM",
	// 		domainID: "172.31.0.1",
	// 		name:     "lisa",
	// 		writeInfo: writeInfo{
	// 			loginUser:   "neo4j",
	// 			loginPasswd: "na",
	// 			configURL:   "http://172.31.86.190:7474/db/data/cypher",
	// 			command:     "CreateNodes",
	// 		},
	// 	},
	// }

	// CreateMultiNodes(commitArray2)
}
