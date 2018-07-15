// 13_闭包捕获外部变量的特点
package main

import (
	"fmt"
)

//定义一个函数返回一个闭包函数
//闭包不关心捕获的变量和常量是否已经超出了作用域,所以只有闭包还在使用它,这些变量就还会存在
func testFunc() func() int {
	var a int
	return func() int {
		a++
		return a * a
	}
}

func main() {
	//	a := 1212
	//	func() {
	//		a = 120
	//		fmt.Println("inside a is ", a)
	//	}()
	//	fmt.Println("outside a is ", a)
	test := testFunc()
	fmt.Println("result is ", test())
	fmt.Println("result is ", test())
	fmt.Println("result is ", test())
}
