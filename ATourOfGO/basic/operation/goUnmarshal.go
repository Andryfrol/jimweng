// refer to website
// https://gobyexample.com/json

package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Page   int           `json:"page"`
	Fruits []interface{} `json:"fruits"`
}

type AutoGenerated struct {
	Columns []string        `json:"columns"`
	Data    [][]interface{} `json:"data"`
}

func main() {
	str := `{
		"page": 1,
		 "fruits": [["apple","banna"], "peach"]
		}`
	res := Response{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)

	strs := `{
		"columns" : [ "id(r)", "n1.domainId", "type(r)", "n2.domainId" ],
		"data" : [ [ 74, "testDomainID3", "take", "testDomainID" ], [ 73, "testDomainID", "belong", "testDomainID2" ], [ 75, "testDomainID2", "follow", "testDomainID3" ], [ 76, "testDomainID", "after", "testDomainID3" ] ]
	  }`

	res2 := AutoGenerated{}
	json.Unmarshal([]byte(strs), &res2)
	fmt.Println(res2.Data[0])
	fmt.Println(res2.Data[0][0])
	fmt.Println(len(res2.Data[0]))

}