package main

import (
	"flag"
	"fmt"
)

var fDebug = flag.Bool("debug", true, "show debug mode help")
var fString = flag.String("name","Jim","default name would be Jim")

type test struct {
	Debug bool
}
func StringInSlice(a string, list []string) bool{
	for _, b := range list{
		if b == a{
			return true
		}
	}
	return false
}

// go run gotest.go -debug=false
func main() {
	flag.Parse()
	fmt.Println(*fDebug)
	if *fDebug {
		fmt.Println("it's right")
	}

	fmt.Println(*fString)

	list :=[]string{"jim","bob","beck"}
	if StringInSlice(*fString, list) {
		fmt.Println("true")
	}

}
