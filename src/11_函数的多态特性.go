// 11_函数的多态特性
package main

import (
	"fmt"
)

//定义一个函数类型
type FuncType func(int, int) int

//定义一个函数
func add(a int, b int) int {
	return a + b
}

//定义一个回调函数
func Calc(a int, b int, testFunc FuncType) (result int) {
	value := testFunc(a, b)
	return value
}

func main() {
	result := Calc(3, 8, add)
	fmt.Println("result is ", result)
}
