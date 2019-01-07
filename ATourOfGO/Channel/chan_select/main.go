package main

import "time"
import "fmt"

func main() {

	// 宣告一個chan c1
	c1 := make(chan string, 1)

	// 在下面的goroutine中將'result 1'送入c1
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	// 使用select來做邏輯判斷，判斷要印出的res
	select {
	case res := <-c1:
		fmt.Println(res)
	// 如果超過1秒就印出timeout 1並且跳出回圈
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}
}
