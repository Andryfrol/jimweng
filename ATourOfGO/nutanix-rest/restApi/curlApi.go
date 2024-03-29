// https://github.com/Tfindelkind/Nutanix_GO_tutorial/blob/master/basic_example/basic.go
package main

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// EncodeCredentials this func is encoding the Username and Password with base64 encoding which is
// required for Nutanix
func EncodeCredentials(username string, password string) string {
	return base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
}

// v0_8 returns the main entry point for the v0.8 Nutanix API
func v0_8(NutanixHost string) string {

	return "https://" + NutanixHost + ":9440/api/nutanix/v0.8/"

}

// v1_0 returns the main entry point for the v1.0 Nutanix API
func v1_0(NutanixHost string) string {

	return "https://" + NutanixHost + ":9440/PrismGateway/services/rest/v1/"

}

// v2_0 returns the main entry point for the v2.0 Nutanix API
func v2_0(NutanixHost string) string {

	return "https://" + NutanixHost + ":9440/PrismGateway/services/rest/v2.0/"

}

// v3_0 returns the main entry point for the v3.0 Nutanix API -> Not GA with AOS 5.0
func v3_0(NutanixHost string) string {

	return "https://" + NutanixHost + ":9440/PrismGateway/services/rest/v3.0/"

}

func main() {

	// PRISM user
	var username = "admin"
	// PRISM user password
	var password = "RyfUA8xC3b}7@3["
	// Nutanix Cluster IP/DNSName CVM IP/DNSName
	var NutanixHost = "10.2.10.119"

	// Ignores certificates which can not be validated
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// create a HTTP client
	var httpClient = http.Client{Transport: tr}

	// create a http Request pointer
	var req *http.Request
	var err error

	// Defines the HTTP Request
	// send a GET to the NUTANIX API and receives the user session_info
	// https://NutanixHost:9440/PrismGateway/services/rest/v1//users/session_info
	// req, _ = http.NewRequest("GET", v1_0(NutanixHost)+"/users/session_info", nil)
	req, _ = http.NewRequest("GET", v1_0(NutanixHost)+"/users/session_info", nil)

	fmt.Println(v1_0(NutanixHost))

	// before the request is send set the HTTP Header key "Authorization" with
	// the value of base64 encoded Username and Password
	req.Header.Set("Authorization", "Basic "+EncodeCredentials(username, password))

	fmt.Printf("----%v\n", req.Header)

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	// Status Code 401 Unauthorized means user+password was not valid
	// https://en.wikipedia.org/wiki/List_of_HTTP_status_codes
	if resp.StatusCode == 401 {
		log.Fatal("Username or password not valid for host: " + NutanixHost)
		os.Exit(1)
	}

	// Response status code 200 should be send if credentials are valid
	// all other could be ignored or handle if needed
	if resp.StatusCode != 200 {
		log.Fatal("Connection to host: " + NutanixHost + " not possible")
		os.Exit(1)
	}

	// read the data from the resp.body into htmlData
	htmlData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// print the response body (htmlData) to give you a feedback
	fmt.Println(string(htmlData))

}
