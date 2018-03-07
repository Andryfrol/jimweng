package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("GET", "https://10.2.10.119:9440/api/nutanix/v2.0/cluster/", nil)
	if err != nil {
		// handle err
	}
	req.SetBasicAuth("admin", "RyfUA8xC3b}7@3[")
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
	for k, v := range pritntVar {
		fmt.Printf("%v__%v\n", k, v)
	}

	defer resp.Body.Close()

}

type clusterStruct struct {
}
