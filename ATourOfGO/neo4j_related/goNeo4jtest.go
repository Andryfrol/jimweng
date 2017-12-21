package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"fmt"
)

// var testString = "neo4j,nodeNum=n1,tag=testTag,domainID=testDomainID, name=testName,link=n1_n2|n3_n1,relation=belong|take,nodeNum=n2,tag=testTag2,domainID=testDomainID2, name=testName2,link=n2_n3,relation=follow,nodeNum=n3,tag=testTag3,domainID=testDomainID3, name=testName3,link=n1_n3,relation=after"

// claim node : would declare properties that would be writed to neo4j
type neo4jMetric struct {
	NAME string `neo4j`
	NodeNum map[string]nodeInfo{}
}

type nodeInfo struct {
	TAG      string `tag`
	DomainID string `domainID`
	NAME	string `name`
	Link     string	`link`
	Relation string `relation`
}

func main() {

	var test :=neo4jMetric{
		
	}

}
