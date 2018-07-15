// 44_测试goroutine
package main

import (
	"fmt"
	"time"
)

func testGoroutine() {
	for {
		fmt.Println("this is new task")
		time.Sleep(time.Second)
	}
}

func main() {
	go testGoroutine()

	for {
		fmt.Println("this is main task")
		time.Sleep(time.Second)
	}

}
