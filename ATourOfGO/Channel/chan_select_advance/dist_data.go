package main

import (
	"fmt"
)

var intArray = []int{1, 2, 3, 4, 5}

func ping(pings chan<- int, msg int) {
	pings <- msg
}

func pong(pings <-chan int, pongs chan<- int) {
	msg := <-pings
	pongs <- msg
}

var counter int

func main() {
	counter = 0
	// var wg = sync.WaitGroup{}
	pings := make(chan int, 5)
	pongs := make(chan int, 5)

	for _, j := range intArray {
		ping(pings, j)
		counter++
	}

	for {
		pong(pings, pongs)
		// defer close(pongs)
		// select {
		// case x := <-pongs:
		fmt.Println(<-pongs)
		// fmt.Println(x)
		// }
		if counter == 1 {
			break
		}
		counter--
	}

}
