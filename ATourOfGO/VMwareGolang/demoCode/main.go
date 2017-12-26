// https://hawksnowlog.blogspot.tw/2017/06/fetch-vapps-with-govmomi.html
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

var (
	envURL         = "https://172.31.17.100/sdk"
	urlDescription = fmt.Sprintf("ESX or vCenter URL [%s]", envURL)
	urlFlag        = flag.String("url", envURL, urlDescription)

	envInsecure         = true
	insecureDescription = fmt.Sprintf("Don't verify the server's certificate chain [%s]", envInsecure)
	insecureFlag        = flag.Bool("insecure", envInsecure, insecureDescription)
)

func main() {
	// vCenter への接続
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	flag.Parse()
	u, err := url.Parse(*urlFlag)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	u.User = url.UserPassword("agent.test", "agent.test")
	c, err := govmomi.NewClient(ctx, u, *insecureFlag)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(c.Client.Client.Version)
	// fmt.Println(c.Client.Client.UserAgent)

	// データセンターの取得
	f := find.NewFinder(c.Client, true)

	// fmt.Println("----")
	// fmt.Println(f)
	// fmt.Println("----")

	dc, err := f.Datacenter(ctx, "DiskProphet")
	// dc, err := f.DefaultDatacenter(ctx)

	// fmt.Println("----")
	// fmt.Println("the diskprophet center return is", dc, err)
	// fmt.Println("----")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	f.SetDatacenter(dc)
	fmt.Println(dc)

	// fmt.Println("-------")
	// fmt.Println(f)
	// fmt.Println("-------")
	// 全 VirtualApp の取得
	// pc := property.DefaultCollector(c.Client)
	// pc := property.Collector(c.Client)
	// fmt.Println("pc value is", &pc)

	// varefs := []types.ManagedObjectReference{}
	// vas, err := f.VirtualAppList(ctx, "*")
	vas, err := f.VirtualMachineList(ctx, "*")
	// fmt.Println("vas is ", vas, err)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// for _, va := range vas {
	// 	// fmt.Println(va.Common.InventoryPath)
	// 	// func (c Common) ObjectName(ctx context.Context) (string, error) {
	// 	// fmt.Println(va.Common.ObjectName(ctx))
	// 	// fmt.Println(va.Common.Name())
	// 	// fmt.Println(va.MarkAsVirtualMachine(ctx).VAppConfig.EntityConfig)
	// 	// var f interface{}
	// 	f, _ := va.QueryConfigTarget(ctx)
	// 	fmt.Printf("", f.Network)
	// 	varefs = append(varefs, va.Reference())
	// }
	f1, _ := vas[0].QueryConfigTarget(ctx)
	for _, t := range f1.Network {
		fmt.Printf("%+v", t)
	}
	// varefs = append(varefs, va.Reference())

	// record VM id
	// fmt.Println(varefs)

	// var vadst []mo.VirtualApp

	// err = pc.Retrieve(ctx, varefs, nil, &vadst)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// for _, va := range vadst {
	// 	fmt.Println(va.VAppConfig.Annotation)
	// 	fmt.Println(va.VAppConfig.EntityConfig) // VApp 内の VM の情報
	// 	fmt.Println(va.Name)
	// }
}
