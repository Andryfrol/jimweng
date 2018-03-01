// https://stackoverflow.com/questions/23558425/how-do-i-get-the-local-ip-address-in-go
package main

import (
	"fmt"
	"net"
	"os"
)

func getIp1() {
	name, err := os.Hostname()
	if err != nil {
		fmt.Printf("Oops: %v\n", err)
		return
	}

	addrs, err := net.LookupHost(name)
	if err != nil {
		fmt.Printf("Oops: %v\n", err)
		return
	}

	for _, a := range addrs {
		fmt.Println(a)
	}
}

func getIp2() {
	ifaces, err := net.Interfaces()
	// handle err
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		// handle err
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			// process IP address
			fmt.Printf("ip address is %v\n", ip)
		}
	}
}

func main() {
	// getIp1()
	getIp2()
}
