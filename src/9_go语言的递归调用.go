// 9_go语言的递归调用
package main

import (
	"fmt"
)

func testRecurse(num int) (result int) {
	var product int
	if num == 1 {
		return 1
	}
	product = num * testRecurse(num-1)
	num--
	return product
}

//用递归来实现数的累加
func testAdd(num int) int {
	//先定义一个返回值
	var result int
	//给出结束条件,防止栈溢出
	if num == 1 {
		return 1
	}
	result = num + testAdd(num-1)
	num--
	return result
}

func main() {
	//	result := testRecurse(4)
	//	fmt.Println("result is ", result)
	result := testAdd(4)
	fmt.Println("result is ", result)
}
