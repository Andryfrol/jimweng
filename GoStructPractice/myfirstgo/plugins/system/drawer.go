package system
// package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main(){
	fmt.Println("This first come out result is")
	fmt.Println("wish to be 2** ",math.Exp2(2))
	var nums int = int(math.Exp2(1))
	var random_num int = rand.Intn(nums)
	fmt.Println(random_num)
	
}
