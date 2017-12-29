package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25/mo"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/find"
)

type Neo4j struct {
	Urls               string
	InsecureSkipVerify bool
}

type nodeInfo struct {
	NodeNum  string
	DomainID string
	Name     string
	Labels   interface{}
	Types    string
}

func vCenterVmName(neo4j Neo4j) map[int]nodeInfo {
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
	viewNewManager := view.NewManager(c.Client)

	fmt.Println(c.PropertyCollector())

	jim, _ := viewNewManager.CreateContainerView(ctx, c.ServiceContent.RootFolder, []string{"HostSystem"}, true)
	if err != nil {
		log.Fatal(err)
	}
	defer jim.Destroy(ctx)

	var hss []mo.HostSystem
	err = jim.Retrieve(ctx, []string{"HostSystem"}, []string{"summary"}, &hss)
	if err != nil {
		log.Fatal(err)
	}

	// Print VMHost
	for _, hs := range hss {
		fmt.Printf("%s\t", hs.Summary.Config.Name)
		// fmt.Println()
	}

	fmt.Println()

	fmt.Println("------------above is host IP---------------")

	f := find.NewFinder(c.Client, true)

	// fmt.Println(f.HostSystemList)

	datacenterList, err := f.DatacenterList(ctx, "*")
	fmt.Println(len(datacenterList))

	// datacenterList(f.DatacenterList(ctx,"*")) would return VMDataCenter
	for i := 0; i < len(datacenterList); i++ {
		fmt.Println(datacenterList[i].ObjectName(ctx))
	}
	fmt.Println("----------above would list vmware VMDataCenter-----------")

	// fmt.Println(datacenterList[1])

	objectNameOfDatacenter, err := datacenterList[1].ObjectName(ctx)

	// fmt.Println(objectNameOfDatacenter)

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

	test, _ := f.VirtualAppList(ctx, "*")

	fmt.Println(test)
	fmt.Println("-------------------------")

	// DatastoreList(ctx,"*") would return VMDatastore -- original default VMDataCenter is "DiskProphet"
	// fmt.Println(f.DatastoreList(ctx, "*"))
	i, _ := f.DatastoreList(ctx, "*")
	for index := 0; index < len(i); index++ {
		objectNameOfDatastores, _ := i[index].ObjectName(ctx)
		fmt.Printf("%s ", objectNameOfDatastores)
		// w, _ := i[index].
		// fmt.Println(w)
	}
	fmt.Println()
	// fmt.Println(i[0].AttachedHosts(ctx))

	fmt.Println("VMDatastore nodes are", len(i))
	fmt.Println("----------above would list vmware VMDataCenter-----------")

	// varefs := []types.ManagedObjectReference{}

	s := make(map[int]nodeInfo, len(vas))
	fmt.Println(len(vas))
	for index, va := range vas {
		keyString := fmt.Sprintf("n%d", index)
		if index == 0 {
			s[index] = nodeInfo{
				NodeNum:  keyString,
				DomainID: va.Name(),
				Name:     va.Name(),
				Types:    "nodes",
				Labels:   "nodes",
			}
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
