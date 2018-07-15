// 57_select语句的使用
package main

import (
	"fmt"
	"time"
)

func main() {
	//select是监听io操作的
	one := make(chan int)
	timeout := make(<-chan time.Time)

	go func() {
		for i := 0; i < 8; i++ {
			if i == 6 {
				one <- 2
			}
			fmt.Println("i is ", i)
			time.Sleep(time.Second)
		}
	}()

	timeout = time.After(5 * time.Second)
	//	timer = time.NewTimer(2 * time.Second)

	select {
	case <-one:
		fmt.Println("channel one has data")
	case <-timeout:
		fmt.Println("channel two has data")
	}
}
