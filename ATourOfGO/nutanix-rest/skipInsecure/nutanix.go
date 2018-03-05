package main

import (
	"bytes"
	"crypto/tls"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Restclient struct {
	ip, username, password, sub_url, method, body, content_type string
}

func (rc *Restclient) init_param(sub_url, method, body string) {
	rc.sub_url = sub_url
	rc.method = method
	rc.body = body
	rc.content_type = "application/json"
	fmt.Println("sub_url:", rc.sub_url, sub_url)
}

func (rc *Restclient) change_content_type(content_type string) {
	rc.content_type = content_type
}

func (rc Restclient) rest_async_api(get_api_name get_fn_name) (status, op string) {
	status, response := rc.rest_client()
	if status != "200 OK" {
		searchPath := []string{"message"}
		message := jsonParser(response, searchPath)
		searchPath = []string{"reason"}
		reason := jsonParser(response, searchPath)
		fmt.Println("Message : ", message)
		fmt.Println("Reason : ", reason)
		return status, response
	}

	search_state_Path := []string{"status", "state"}
	api_state := jsonParser(response, search_state_Path)
	if api_state != "kComplete" || api_state != "kError" {
		search_uuid_Path := []string{"metadata", "uuid"}
		uuid := jsonParser(response, search_uuid_Path)

		//probe async api status
		for cnt := 0; cnt < 10; cnt++ {
			time.Sleep(5 * time.Second)

			status, response := get_api_name(rc, uuid)
			if status == "200 OK" {
				fmt.Println("Response: ", response)

				api_state = jsonParser(response, search_state_Path)
				if api_state == "kComplete" {
					return status, response
				}
			} else {
				fmt.Println("Failed to get details")
				return status, response
			}
		}
	}

	return status, response
}

func (rc Restclient) rest_client() (status, op string) {
	url := "https://" + rc.ip + ":9440/api/nutanix/v3/" + rc.sub_url
	fmt.Println("url: ", url, rc.method)
	jsonBytes := []byte(rc.body)

	req, err := http.NewRequest(rc.method, url, bytes.NewBuffer(jsonBytes))
	req.Header.Set("Authorization", "Basic "+b64.StdEncoding.EncodeToString([]byte(rc.username+":"+rc.password)))
	req.Header.Set("Content-Type", rc.content_type)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("RestClient failed for url: ", url)
		return "503 Connection failed", "Restclient failed."
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	response := string(body)

	if resp.Status != "200 OK" {
		searchPath := []string{"message"}
		message := jsonParser(response, searchPath)
		searchPath = []string{"reason"}
		reason := jsonParser(response, searchPath)
		fmt.Println("Message : ", message)
		fmt.Println("Reason : ", reason)
		return status, response
	}

	return resp.Status, response
}

type list_fn_name func(rc Restclient) (status, result string)

type get_fn_name func(rc Restclient, uuid string) (status, result string)

func main() {
	var cluster_uuid string
	// cluster_uuid = "00054fd5-6de4-8422-74c5-782bcb637d0e"

	var status, response string

	if len(os.Args) != 4 {
		fmt.Println("Usage: ./filename <server-ip-address>  <login-username>  <login-password>")
		return
	}

	rc := Restclient{ip: os.Args[1], username: os.Args[2], password: os.Args[3]}

	//List all cluster
	status, response = cluster_list(rc)
	if status == "200 OK" {
		fmt.Println("List all cluster response:", response)

		b := []byte(response)
		var jsontype map[string][]map[string]interface{}
		json.Unmarshal(b, &jsontype)
		for i := range jsontype["entities"] {
			item := jsontype["entities"][i]
			jsonType := item["metadata"].(map[string]interface{})
			cluster_uuid = jsonType["uuid"].(string)
		}
	} else {
		fmt.Println("Failed to list all the clusters")
		return
	}

	//Get cluster
	status, response = cluster_get(rc, cluster_uuid)
	if status == "200 OK" {
		fmt.Println("Got cluster details successfully")
		fmt.Println("Response:", response)
	} else {
		fmt.Println("failed to get cluster details")
		return
	}

}

//This function will give list of cluster
func cluster_list(rc Restclient) (status, result string) {
	body := `{ "kind": "cluster",
                   "offset" : 0,
                   "length" : 0 }`

	rc.init_param("clusters/list", "POST", body)
	status, response := rc.rest_client()
	return status, response
}

//This function will return the cluster definition by the given uuid
func cluster_get(rc Restclient, cluster_uuid string) (status, result string) {
	sub_url := "clusters/" + cluster_uuid
	body := ""
	rc.init_param(sub_url, "GET", body)
	status, response := rc.rest_client()
	return status, response
}

//This function will parse json data into golang
func jsonParser(jsonstring string, searchPath []string) (output string) {
	l := len(searchPath)
	j := (l - 1)
	var parsedValue string
	b := []byte(jsonstring)
	var jsontype interface{}
	err := json.Unmarshal(b, &jsontype)
	if err != nil {
		fmt.Println("%s Failed to convert json to go", err)
	}

	for i := 0; i < (l - 1); i++ {
		m := jsontype.(map[string]interface{})
		jsontype = m[searchPath[i]].(map[string]interface{})
		//fmt.Println("UUID = ", searchPath[i], jsontype)
	}

	m := jsontype.(map[string]interface{})
	for k, v := range m {
		if k == searchPath[j] {
			//fmt.Println(v)
			parsedValue = v.(string)
		}
	}
	return parsedValue
}
