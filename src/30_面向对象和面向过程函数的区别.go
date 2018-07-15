// 30_面向对象和面向过程函数的区别
package main

import (
	"fmt"
)

func testFunc(a, b int) (result int) {
	result = a + b
	return
}

type long int

func (temp long) testFunc2(num long) (result long) {
	result = temp + num
	return
}

func main() {
	var num long = 23
	result := num.testFunc2(90)
	fmt.Println("result is ", result)
}
