package main

import (
	"context"
	"flag"
	"fmt"
	"net/url"
	"os"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/find"
)

type Neo4j struct {
	Urls               string
	InsecureSkipVerify bool
}

type nodeInfo struct {
	domainId string
	name     string
	labels   int
}

func vCenterVmName(neo4j Neo4j) map[string]nodeInfo {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	flag.Parse()
	u, err := url.Parse(neo4j.Urls)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	u.User = url.UserPassword("agent.test", "agent.test")
	c, err := govmomi.NewClient(ctx, u, neo4j.InsecureSkipVerify)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	f := find.NewFinder(c.Client, true)

	datacenterList, err := f.DatacenterList(ctx, "*")
	objectNameOfDatacenter, err := datacenterList[1].ObjectName(ctx)

	fmt.Println(objectNameOfDatacenter)

	dc, err := f.Datacenter(ctx, objectNameOfDatacenter)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	f.SetDatacenter(dc)
	vas, err := f.VirtualMachineList(ctx, "*")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(f.DatastoreList(ctx, "*"))

	// varefs := []types.ManagedObjectReference{}

	s := make(map[string]nodeInfo, len(vas))
	for index, va := range vas {
		keyString := fmt.Sprintf("n%d", index)
		if index == 0 {
			s[keyString] = nodeInfo{
				domainId: va.Name(),
				name:     va.Name(),
				labels:   index,
			}
			// t, _ := va.Device(ctx)
			// l := t.ChildDisk()
			// fmt.Println(l)
		} else {
			continue
		}
	}

	return s
}

func main() {
	neo4jTest := Neo4j{
		Urls:               "https://172.31.17.100/sdk",
		InsecureSkipVerify: true,
	}
	// fmt.Printf("%v", vCenterVmName(neo4jTest))
	fmt.Println(vCenterVmName(neo4jTest))
}

// relationship terms
var (
	VMDataCenterLabelName       = "VMDataCenter"
	VMClusterLabelName          = "VMClusterCenter"
	VMHostLabelName             = "VMHost"
	VMVSanClusterLabelName      = "VMVSanCluster"
	VMVSanDiskGroupLabelName    = "VMVSanDiskGroup"
	VMVSanCacheDiskLabelName    = "VMVSanCacheDisk"
	VMVSanCapacityDiskLabelName = "VMVSanCapacityDisk"
	VMDatastoreLabelName        = "VMDatastore"
	VMDiskLabelName             = "VMDisk"
	VMVirtualMachinesLabelName  = "VMVirtualMachine"
	VMSnapshotLabelName         = "VMSnapshot"

	VMDataCenterContainsVMClusterRelationName        = "VmDataCenterContainsVmCluster"
	VMDataCenterContainsVMVSanClusterRelationName    = "VmDataCenterContainsVSanCluster"
	VMVSanClusterContainsVMVSanDiskGroupRelationName = "VSanClusterContainsVSanDiskGroup"
	VMClusterContainsVMHostRelationName              = "VmClusterContainsVmHost"
	VMClusterContainsVMDatastoreRelationName         = "VmClusterContainsVmDatastore"
	VMHostContainsVMDiskRelationName                 = "VmHostContainsVmDisk"
	VMHostHasVMDiskGroupRelationName                 = "VmHostHasVmDiskGroup"
	VMVsanDatastoreContainsVMDiskGroupRelationName   = "VsanDatastoreContainsVmDiskGroup"

	VMVSanDiskGroupHasCacheVMDiskRelationName    = "VSanDiskGroupHasCacheVmDisk"
	VMVSanDiskGroupHasCapacityVMDiskRelationName = "VSanDiskGroupHasCapacityVmDisk"
	VMHostHostsVMVirtualMachineRelationName      = "VmHostHostsVmVirtualMachine"
	VMDatastoreComposesOfVMDiskRelationName      = "VmDatastoreComposesOfVmDisk"
	VMHostHasVMDatastoreRelationName             = "VmHostHasVmDatastore"
	VMVirtualMachineUsesVMDatastoreRelationName  = "VmVirtualMachineUsesVmDatastore"
	VMVirtualMachineTakesVMSnapshotRelationName  = "VmVirtualMachineTakesVmSnapshot"
)
