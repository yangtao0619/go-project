// 45_gosched的使用
package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		for i := 0; i < 2; i++ {
			fmt.Println("go")
		}
	}()

	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println("hello")
	}
}
