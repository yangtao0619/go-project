// 36_panic函数的使用
package main

import (
	"fmt"
)

func testFunc1() {

	fmt.Println("aaaaaaaaaa")
}

func testFunc2() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	panic("process is crashed!")
}

func testFunc3() {
	fmt.Println("ccccccccccc")
}
func main() {
	testFunc1()
	testFunc2()
	testFunc3()
}
