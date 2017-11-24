package main

import (
	"flag"
	"fmt"
)

var fDebug = flag.Bool("debug", true, "show debug mode help")

type test struct {
	Debug bool
}

// go run gotest.go -debug=false
func main() {
	flag.Parse()
	fmt.Println(*fDebug)
	if *fDebug {
		fmt.Println("it's right")
	}
}
