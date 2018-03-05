package main

import (
	"fmt"
	"nutanix"
)

func main() {

	nutanixConf := nutanix.NewConfiguration()

	nutanixConf.Username = "admin"
	nutanixConf.Password = "RyfUA8xC3b}7@3["
	// nutanixConf.Password = "nutanix/4u"
	nutanixConf.BasePath = "https://10.2.10.119:9440/api/nutanix/v3/"

	// nutanixConf.APIClient.

	// create an instance of the API class
	// apiInstance := nutanix.ClustersApi{Configuration: *nutanixConf}
	// testResp := nutanix.ApiRequest
	apiInstance := nutanix.VmsApi{*nutanixConf}
	// apiInstance.Configuration.APIClient.CallAPI().Request

	// apiInstance.Configuration.APIClient.CallAPI().Request.SetAuthToken()

	fmt.Printf("%v\n", apiInstance)

	body := nutanix.VmIntentInput{}

	nResp, apiResp, err := apiInstance.CreateVm(body)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("%v\n", nResp)
	fmt.Printf("%v\n", apiResp)
}
