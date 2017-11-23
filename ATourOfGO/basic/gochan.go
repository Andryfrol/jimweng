// https://mikespook.com/2013/05/%E7%BF%BB%E8%AF%91%E7%BB%9D%E5%A6%99%E7%9A%84-channel/

package main

import (
	"fmt"
	"sync"
	"time"
)

func example1() {
	ch := make(chan bool, 2)
	ch <- true
	ch <- true
	close(ch)

	for i := 0; i < cap(ch)+1; i++ {
		v, ok := <-ch
		fmt.Println(v, ok)
	}
}

func example2() {
	finish := make(chan bool)
	var done sync.WaitGroup
	done.Add(1)
	go func() {
		select {
		case <-time.After(1 * time.Hour):
		case <-finish:
		}
		done.Done()
	}()
	t0 := time.Now()
	finish <- true
	done.Wait()
	fmt.Printf("Waited %v for goroutine to stop\n", time.Since(t0))
}

func main() {
	example2()
}
