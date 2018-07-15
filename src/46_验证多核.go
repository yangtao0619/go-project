// 46_验证多核
package main

import (
	"fmt"
	"runtime"
)

func main() {
	n := runtime.GOMAXPROCS(4)
	fmt.Println("n is ", n)
	for {
		go fmt.Print("1")

		fmt.Print("0")
	}
}
