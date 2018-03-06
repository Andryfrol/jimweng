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
	// clusterApiInstance := nutanix.ClustersApi{Configuration: *nutanixConf}
	// body := nutanix.ClusterListMetadata{}

	// clusterResp, clusterApiResp, err := clusterApiInstance.ListCluster(body)
	// if err != nil {
	// 	fmt.Printf("%v\n", err)
	// }
	// fmt.Printf("\tclusterResp %v\n", clusterResp)
	// fmt.Printf("\n\tclusterApiResp %v\n", clusterApiResp)

	vmApiInstance := nutanix.VmsApi{Configuration: *nutanixConf}
	body2 := nutanix.VmListMetadata{}

	vmResp, vmApiResp, err := vmApiInstance.ListVm(body2)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	// fmt.Printf("---------%v\n", vmResp)

	// fmt.Printf("\tvmResp 1 %v\n", vmResp.Entities)
	for _, vmir := range vmResp.Entities {
		// fmt.Printf("0. %v\n vmdir is", vmir)
		// fmt.Printf("1. %v\n", vmir.Metadata)
		fmt.Printf("1. vm_Name %v\n", vmir.Metadata.Name) //wrong vm name
		fmt.Printf("1-2. vm_Uuid %v\n", vmir.Metadata.Uuid)
		fmt.Printf("3-2. vm_Status_Name %v\n", vmir.Status.Name)
		// fmt.Printf("1-3. vm_Uuid %v\n", vmir.Metadata)

		// fmt.Printf("2. %v\n", vmir.Spec)
		// fmt.Printf("3-1. vm_ClusterReference %v\n", vmir.Status.ClusterReference)
		// fmt.Printf("3-2. vm_MessageList %v\n", vmir.Status.MessageList)

		// fmt.Printf("3-3. vm_Status_ %v\n", vmir.Status)

		fmt.Printf("3-4. vm_Status_MemorySizeMib %v\n", vmir.Status.Resources.MemorySizeMib)

		fmt.Printf("3-5. vm_Status_Resource_NumVcpu %v\n", vmir.Status.Resources.NumVcpusPerSocket)
		fmt.Printf("3-6. vm_Status_Resource_PowerStateMechanism%v\n", vmir.Status.Resources.PowerStateMechanism)

		for _, disk := range vmir.Status.Resources.DiskList {
			fmt.Printf("\t disk: %v\n", disk)
		}

		for _, gpu := range vmir.Status.Resources.GpuList {
			fmt.Printf("\t gpu: %v\n", gpu)
		}
	}

	fmt.Printf("\n\n\n\n\n\tvmResp %v\n", vmResp.Metadata)
	fmt.Printf("\n\tvmApiResp %v\n", vmApiResp)

}
