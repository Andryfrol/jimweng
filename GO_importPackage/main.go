package main

import (
	"fmt"

	"github.com/goPractice/GO_importPackage/goexample"
	"github.com/goPractice/GO_importPackage/string"
)

func main() {
	goexample.Hi()
	goexample.Hello()
	var test = string.Reverse("Hello")
	fmt.Println(test)
}
