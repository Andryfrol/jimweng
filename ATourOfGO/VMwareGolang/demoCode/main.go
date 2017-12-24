// Golang context refernces
// https://deepzz.com/post/golang-context-package-notes.html
// http://dev.twsiyuan.com/2017/09/golang-iterator-channel-implementation.html
// http://www.nljb.net/default/Golang之Context的使用/
// https://github.com/vmware/govmomi/blob/master/client.go

package main

import (
	"context"
	"fmt"
	"net/url"

	"github.com/vmware/govmomi/session"
	"github.com/vmware/govmomi/vim25"
)

type Client struct {
	*vim25.Client

	SessionManager *session.Manager
}

func main() {
	var ctx *context.Context
	fmt.Println(ctx)
	test := "http://172.31.86.190:7474/db/data"
	testurl, _ := url.Parse(test)
	fmt.Println(testurl)
	// c, err := govmomi.NewClientWithCertificate(ctx, "", "")
	// fmt.Println(govmomi.IsVC(*c))
}
