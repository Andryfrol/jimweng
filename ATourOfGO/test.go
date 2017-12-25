// Basic example of getting disk.usage performance information from ESX hosts
// using govmomi. This code does not include any error checking right now on
// purpose.
//
// This program relies on an environment variable called VC_URL, which would be
// the URL to the vCenter SDK. For example:
//
//      VC_URL=https://user:pass@vcenter-ip/sdk go run this_code.go
//
package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"reflect"

	"golang.org/x/net/context"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
)

func main() {
	// url, _ := url.Parse(os.Getenv("VC_URL"))
	// https://user:pass@vcenter-ip/sdk
	// https://agent.test:agent.test@172.31.17.100/sdk
	// url, _ := url.Parse("https://agent.test:agent.test@172.31.17.100/sdk")
	url, _ := url.Parse("https://172.31.17.100/sdk")

	fmt.Println(os.Getenv)
	fmt.Println(reflect.TypeOf(url))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, _ := govmomi.NewClient(ctx, url, true)
	fmt.Println(client)

	finder := find.NewFinder(client.Client, true)
	fmt.Println()

	dc, _ := finder.DefaultDatacenter(ctx)
	finder.SetDatacenter(dc)

	hosts, _ := finder.HostSystemList(ctx, "*")

	// disk.usage
	metricId := types.PerfMetricId{CounterId: 125, Instance: "*"}

	for _, host := range hosts {
		querySpec := types.PerfQuerySpec{
			Entity:     host.Reference(),
			MaxSample:  1,
			MetricId:   []types.PerfMetricId{metricId},
			IntervalId: 20,
		}
		query := types.QueryPerf{
			This:      *client.ServiceContent.PerfManager,
			QuerySpec: []types.PerfQuerySpec{querySpec},
		}

		res, _ := methods.QueryPerf(ctx, client, &query)
		muck := res.Returnval[0]
		log.Printf("Response: (%T) %+v", muck, muck)

		// causes "muck.SampleInfo undefined (type types.BasePerfEntityMetricBase has no field or method SampleInfo)"
		// log.Printf("SampleInfo: %+v", muck.SampleInfo)
	}
}
