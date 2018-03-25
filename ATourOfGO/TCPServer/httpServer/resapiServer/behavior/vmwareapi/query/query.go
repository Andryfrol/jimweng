package query

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/vmware/govmomi/find"

	"github.com/vmware/govmomi"
)

type Query struct {
	EnvURL      string
	EnvUserName string
	EnvPassword string
}

func (q *Query) ListDatacenter(req *http.Request) string {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	u, err := url.Parse(q.EnvURL)
	if err != nil {
		return fmt.Sprintf("%s\n", err)
	}

	u.User = url.UserPassword("agent.test", "agent.test")
	c, err := govmomi.NewClient(ctx, u, true)
	if err != nil {
		return fmt.Sprintf("%s\n", err)
	}

	f := find.NewFinder(c.Client, true)

	dc, err := f.Datacenter(ctx, "DiskProphet")
	if err != nil {
		return fmt.Sprintf("%s\n", err)
	}

	return fmt.Sprintf("printout datacenter %s\n", dc)
}
