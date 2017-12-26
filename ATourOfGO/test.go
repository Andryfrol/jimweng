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

func vCenterVmName(neo4j Neo4j) interface{} {
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

	dc, err := f.Datacenter(ctx, "DiskProphet")

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

	// varefs := []types.ManagedObjectReference{}
	s := make(map[string]interface{}, len(vas))

	for _, va := range vas {
		// fmt.Printf("%d\n", index)
		// fmt.Printf("%s\n", va.Common.Name())
		// fmt.Println(va.QueryConfigTarget.Name)
		// fmt.Println(reflect.TypeOf(va.Common.Name()))
		s[va.Common.Name()] = va.Common.Name()
		// varefs = append(varefs, va.Reference())
		// varefs = append()
	}
	return s
}

func main() {
	neo4jTest := Neo4j{
		Urls:               "https://172.31.17.100/sdk",
		InsecureSkipVerify: true,
	}
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
