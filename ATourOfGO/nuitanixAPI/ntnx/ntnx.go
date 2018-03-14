package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/goPractice/ATourOfGO/nuitanixAPI/ntnx/cluster"
	"github.com/goPractice/ATourOfGO/nuitanixAPI/ntnx/disks"
	"github.com/goPractice/ATourOfGO/nuitanixAPI/ntnx/hosts"
	"github.com/goPractice/ATourOfGO/nuitanixAPI/ntnx/storagecontainers"
	"github.com/goPractice/ATourOfGO/nuitanixAPI/ntnx/storagepool"
	"github.com/goPractice/ATourOfGO/nuitanixAPI/ntnx/virtualmachine"
	"github.com/goPractice/ATourOfGO/nuitanixAPI/ntnx/volumegroups"
)

// basic information
var (
	Username = "admin"
	Password = "RyfUA8xC3b}7@3["
	HostIP   = "10.2.10.119"
)

// NtnxClient for nutanix client information
type NtnxClient struct {
	Username     string
	Password     string
	HostIP       string
	Client       *http.Client
	Measurements map[endPoint]exploerAPI
}

// v2_0 returns the main entry point for the v2.0 Nutanix API
func v2_0(NutanixHost string) string {
	return "https://" + NutanixHost + ":9440/PrismGateway/services/rest/v2.0/"
}

// v1_0 returns the main entry point for the v1.0 Nutanix API
func v1_0(NutanixHost string) string {
	return "https://" + NutanixHost + ":9440/PrismGateway/services/rest/v1/"
}

func (ntnxc *NtnxClient) queryNutanixInfor(queryEndPt string) (*[]byte, error) {

	req, err := http.NewRequest("GET", queryEndPt, nil)
	if err != nil {
		// encount error while make req to query nutanix info
		log.Printf("E! encount error while make req to query nutanix info")
		return nil, err
	}

	req.SetBasicAuth(ntnxc.Username, ntnxc.Password)
	req.Header.Set("Accept", "application/json")

	resp, err := ntnxc.Client.Do(req)
	if err != nil {
		// encount error while do req to query nutanix info
		log.Printf("E! encount error while do req to query nutanix info")
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// encount error while do io read
		log.Printf("E! encount error while do io read")
		return nil, err
	}

	return &bodyBytes, nil
}

type endPoint string

const (
	ntnxCluster           = endPoint("cluster")
	ntnxStoragepool       = endPoint("storage_pools")
	ntnxDisks             = endPoint("disks")
	ntnxHosts             = endPoint("hosts")
	ntnxStorageContainers = endPoint("storage_containers")
	ntnxVirtualMachine    = endPoint("vms")
	ntnxVolumeGroups      = endPoint("volume_groups")
)

type exploerAPI string

const (
	ntnxV1 = exploerAPI("v1")
	ntnxV2 = exploerAPI("v2")
)

func main() {
	measurements2 := map[endPoint]exploerAPI{
		ntnxCluster:           ntnxV2,
		ntnxStoragepool:       ntnxV1,
		ntnxDisks:             ntnxV2,
		ntnxHosts:             ntnxV2,
		ntnxStorageContainers: ntnxV2,
		ntnxVirtualMachine:    ntnxV2,
		ntnxVolumeGroups:      ntnxV2,
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	ntnxClient := NtnxClient{
		Username: Username,
		Password: Password,
		HostIP:   HostIP,
		Client: &http.Client{
			Transport: tr,
		},
		Measurements: measurements2,
	}
	var jsonResp interface{}
	for m, v := range ntnxClient.Measurements {
		if v == ntnxV2 {
			queryString := v2_0(ntnxClient.HostIP) + fmt.Sprintf("%s", m)
			bodyBytes, _ := ntnxClient.queryNutanixInfor(queryString)

			switch m {
			case ntnxCluster:
				jsonResp = new(cluster.NtnxCluster)
			case ntnxDisks:
				jsonResp = new(disks.NtnxDisks)
			case ntnxHosts:
				jsonResp = new(hosts.NtnxHosts)
			case ntnxStorageContainers:
				jsonResp = new(storagecontainers.NtnxStorageContainers)
			case ntnxVirtualMachine:
				jsonResp = new(virtualmachine.NtnxVirtualMachine)
			case ntnxVolumeGroups:
				jsonResp = new(volumegroups.NtnxVolumeGroups)
			}
			json.Unmarshal(*bodyBytes, &jsonResp)

			// print and check
			if m == ntnxVirtualMachine {
				fmt.Printf("%v____%v\n", v, queryString)
				fmt.Printf("%v\n", jsonResp)
			}

		}
		if v == ntnxV1 {
			queryString := v1_0(ntnxClient.HostIP) + fmt.Sprintf("%s", m)
			bodyBytes, _ := ntnxClient.queryNutanixInfor(queryString)

			switch m {
			case ntnxStoragepool:
				jsonResp = new(storagepool.StoragePool)
			}
			json.Unmarshal(*bodyBytes, &jsonResp)
		}

	}
}
