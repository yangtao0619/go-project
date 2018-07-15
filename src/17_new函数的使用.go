// 17_new函数的使用
package main

import (
	"fmt"
)

func main() {
	//new函数可以返回一个指针的变量
	p := new(int)
	*p = 999
	fmt.Println("*p is ", *p)
}
