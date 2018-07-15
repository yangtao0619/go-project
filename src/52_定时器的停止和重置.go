// 52_定时器的停止和重置
package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(2 * time.Second)

	go func() {
		for {
			<-timer.C
			fmt.Println("定时时间到")
		}
	}()

	timer.Stop()

	for {

	}
}
