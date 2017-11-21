package main

import (
	"fmt"
)

func main() {
	var m map[string]string

	m = make(map[string]string)
	m["Jim"] = "weng"

	fmt.Println(m)

}
