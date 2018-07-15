// 16_指针的基本操作
package main

import (
	"fmt"
)

func main() {
	//定义一个指针变量
	var p *int
	var a = 10
	fmt.Println("a is ", a)
	p = &a
	*p = 999
	fmt.Println("after a is ", a)
	new()
}
