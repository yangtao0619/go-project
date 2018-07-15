// 48_没有缓冲区的管道,没有缓冲区的管道
package main

import (
	"fmt"
)

func main() {
	//创建一个无缓冲区的管道
	channel := make(chan string)

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("task one")
			channel <- "hello"
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			<-channel
			fmt.Println("task two")
		}
	}()
	for {
	}
}
