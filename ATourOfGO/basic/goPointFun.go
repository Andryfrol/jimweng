// https://gobyexample.com/pointers
package main

import (
	"fmt"
)

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	a := 1
	fmt.Println("initial:", a)

	zeroval(a)
	fmt.Println("zeroval:", a)

	zeroptr(&a)
	fmt.Println("zeroval:", a)

	fmt.Println("zeroptr:", &a)
}
