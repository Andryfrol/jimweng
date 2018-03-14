package nutanix

import (
	"fmt"
	ntx "nutanix"
)

// nutanixConf.Username = "admin"
// nutanixConf.Password = "RyfUA8xC3b}7@3["
// nutanixConf.BasePath = "https://10.2.10.119:9440/api/nutanix/v3/"

// Nutanix LoginInfo
type Nutanix struct {
	NutanixClusters []*NtnxCluster
	Configuration   ntx.Configuration
}

func NewNutanixConfig(username string, password string, ip string) *Nutanix {

	dpNutanixCluster := new(Nutanix)
	nutanixConf := ntx.NewConfiguration()
	nutanixConf.BasePath = "https://" + ip + ":9440/api/nutanix/v3/"
	nutanixConf.Username = username
	nutanixConf.Password = password
	dpNutanixCluster.Configuration = *nutanixConf

	return dpNutanixCluster
}

type NtnxCluster struct {
	BasicInfo
	Hosts []*NtnxHost
}

type NtnxHost struct {
	BasicInfo
	VMs        []*NtnxVirtualMachine
	Datastores []*NtnxDatastore
}

type NtnxVirtualMachine struct {
	BasicInfo
	Memory  string
	Volumne string
}

type NtnxDatastore struct {
	BasicInfo
	Disks []*NtnxDisk
}

type NtnxDisk struct {
	BasicInfo
}

type BasicInfo struct {
	Name string
	Uuid string
}

// Disconnector
func (ntnx *Nutanix) DisConnectNutanix() error {
	return nil
}

// create Nutanix clusters
func (ntnx *Nutanix) CreateNutanixCluster() error {

	clusterApiInstance := ntx.ClustersApi{Configuration: ntnx.Configuration}
	body := ntx.ClusterListMetadata{}

	clusterResp, clusterApiResp, _ := clusterApiInstance.ListCluster(body)
	if clusterApiResp.Status != "200 OK" {
		return fmt.Errorf("fail to get cluster info")
	}
	//// lack of cluster uuid info
	for _, cl := range clusterResp.Entities {
		ntnxcl := &NtnxCluster{BasicInfo: BasicInfo{Name: cl.Status.Name, Uuid: cl.Status.Resources.Config.Build.CommitId}}

		ntnx.NutanixClusters = append(ntnx.NutanixClusters, ntnxcl)
	}

	return nil
}

// nutanix cluster append hosts
func (ntnxcl *NtnxCluster) AppendHts(ht *NtnxHost) error {
	if ht == nil {
		return fmt.Errorf("append a nil hosts in nutanix cluster")
	} else {
		ntnxcl.Hosts = append(ntnxcl.Hosts, ht)
	}
	return nil
}

// append datastores
func (ntnxht *NtnxHost) AppendDss(ds *NtnxDatastore) error {
	if ds == nil {
		return fmt.Errorf("append a nil datastore to nutanix cluster")
	} else {
		ntnxht.Datastores = append(ntnxht.Datastores, ds)
	}
	return nil
}

// append clusters
func (ntnxht *NtnxHost) AppendVMs(vm *NtnxVirtualMachine) error {
	if vm == nil {
		return fmt.Errorf("append a nil virtual machine to nutanix cluster")
	} else {
		ntnxht.VMs = append(ntnxht.VMs, vm)
	}
	return nil
}

// Nutanix connect Cluster. Cluster includes hosts ,and datastores.
// Datastore includes disks
// hosts includes vms
// vm includes information like cpu/mem/diskSize
func (ntnx *Nutanix) GetTopology() error {
	err := ntnx.CreateNutanixCluster()
	if err != nil {
		return err
	}
	// not find the connection between cluster and host yet.
	// since there exists only one cluster, appending all hosts
	// on it. would improve in the future
	apiHostInstance := ntx.HostsApi{Configuration: ntnx.Configuration}
	hostBody := ntx.HostListMetadata{}

	hostResp, _, _ := apiHostInstance.ListHost(hostBody)
	for _, hte := range hostResp.Entities {
		host := NtnxHost{BasicInfo: BasicInfo{Name: hte.Status.Name, Uuid: hte.Status.Resources.SerialNumber}}
		ntnx.NutanixClusters[0].AppendHts(&host)
	}

	apiVmInstance := ntx.VmsApi{Configuration: ntnx.Configuration}
	vmBody := ntx.VmListMetadata{}

	vmResp, _, _ := apiVmInstance.ListVm(vmBody)
	for _, vme := range vmResp.Entities {
		ntnx.findHost(&vme, vme.Status.Resources.HostReference.Uuid)
	}

	ntnx.dumpData()

	return nil
}

func (ntnx *Nutanix) findHost(vmInfo *ntx.VmIntentResource, vmConnectHostUuid string) {
	for _, cl := range ntnx.NutanixClusters {
		for _, ht := range cl.Hosts {
			if vmConnectHostUuid == ht.Uuid {
				vm := &NtnxVirtualMachine{BasicInfo: BasicInfo{Name: vmInfo.Status.Name, Uuid: vmInfo.Metadata.Uuid}}
				ht.AppendVMs(vm)
			}
		}
	}
}

func (ntnx *Nutanix) dumpData() {
	// dump data
	for i, cl := range ntnx.NutanixClusters {
		fmt.Printf("the %d_th cluster name/uuid is %v\t%v\n", i, cl.Name, cl.Uuid)
		for j, ht := range cl.Hosts {
			fmt.Printf("\tthe %d_th host name/uuid is %v\t%v\n", j, ht.Name, ht.Uuid)
			for k, vm := range ht.VMs {
				fmt.Printf("\t\tthe %d_th vm name/uuid is %v\t%v\n", k, vm.Name, vm.Uuid)
			}
		}
	}
}
