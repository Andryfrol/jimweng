// refer to website
// https://gobyexample.com/json

package main

import (
	"encoding/json"
	"fmt"
	"reflect"
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
	fmt.Println(res.Fruits[0].([]interface{})[0])
	test :=res.Fruits[0].([]interface{})[0]
	fmt.Println(reflect.TypeOf(test))
	fmt.Println(test.(string))

	strs :=`{"name" :"data"}`
	var f map[string]interface{}

	json.Unmarshal([]byte(strs), &f)
	fmt.Println(f)
	// fmt.Println(testf["name"])
	// fmt.Println(f.(map[string]interface{})["name"])
	// strs := `{
	// 	"columns" : [ "id(r)", "n1.domainId", "type(r)", "n2.domainId" ],
	// 	"data" : [ [ 74, "testDomainID3", "take", "testDomainID" ], [ 73, "testDomainID", "belong", "testDomainID2" ], [ 75, "testDomainID2", "follow", "testDomainID3" ], [ 76, "testDomainID", "after", "testDomainID3" ] ]
	//   }`

	// res2 := AutoGenerated{}
	// json.Unmarshal([]byte(strs), &res2)
	// fmt.Println(res2.Data[0])
	// fmt.Println(res2.Data[0][0])
	// fmt.Println(len(res2.Data[0]))

}
