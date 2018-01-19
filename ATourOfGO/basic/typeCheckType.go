package main

import (
	"fmt"
)

// parse node info
type NodeInfo struct {
	tag      string
	label    string
	domainId string
	name     string
	vcsa     string
}

type test struct {
	A string
}

type test1 struct {
	B int
}

func genNode(a interface{}) *NodeInfo {

	switch a.(type) {
	case test:
		fmt.Println("string")
	default:
		fmt.Println("not string")
	}

	return nil
}

func main() {
	a := test{
		A: "string",
	}
	b := test1{
		B: 123,
	}
	genNode(a)
	genNode(b)

	fmt.Print("This is example for a.(type)")
}
