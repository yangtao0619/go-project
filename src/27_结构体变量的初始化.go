// 27_结构体变量的初始化
package main

import (
	"fmt"
)

type Student struct {
	age   int
	name  string
	addr  string
	score int
}

func main() {
	//结构体变量的初始化一定要全部赋值
	s1 := Student{12, "mike", "上海", 87}
	fmt.Println("student is ", s1)

	//还可以指定某一个变量进行赋值
	s2 := Student{age: 23, score: 89}
	fmt.Println("student2 is ", s2)
}
