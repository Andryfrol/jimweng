package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	Username = "admin"
	Password = "RyfUA8xC3b}7@3["
	hostIp   = "https://10.2.10.119:9440/api/nutanix/v2.0/clusters/"
)

func main() {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("GET", hostIp, nil)
	if err != nil {
		// handle err
	}
	req.SetBasicAuth(Username, Password)
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		// handle err
	}

	// if needed f would be replaced as some struct
	var f map[string]interface{}
	jsonResp := f
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &jsonResp)
	pritntVar := jsonResp //.(map[string]interface{})
	for i, v := range pritntVar {
		switch v.(type) {
		case map[string]interface{}:
			spanStruct(v)
		default:
			fmt.Printf("%v__%v\n", i, v)
		}
	}

	defer resp.Body.Close()

}

func spanStruct(t interface{}) {
	switch t.(type) {
	case map[string]interface{}:
		for i, v := range t.(map[string]interface{}) {
			fmt.Printf("%v__%v\n", i, v)
			spanStruct(v)
		}
	case []map[string]interface{}:
		for i, v1 := range t.([]map[string]interface{}) {
			fmt.Printf("index %v\t\t", i)
			for j, v := range v1 {
				fmt.Printf("%v__%v\n", j, v)
				spanStruct(v)
			}
		}
	default:
		// fmt.Printf("%v\n", t)
	}
}
