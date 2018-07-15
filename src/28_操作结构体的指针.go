// 28_操作结构体的指针
package main

import (
	"fmt"
)

type Student struct {
	age   int
	name  string
	score int
}

//使用指针改变结构体里面的值
func changeAge(p *Student) {
	p.age = 27
}

func main() {
	//定义一个结构体指针
	//	var s1 *Student = &Student{23, "John", 89}
	//	fmt.Println("s1 is ", s1)

	//	s2 := &Student{29, "Hellen", 90}
	//	fmt.Println("*s2 is ", *s2)
	var s *Student = &Student{23, "Mike", 87}
	fmt.Println("s is ", s)
	changeAge(s)
	fmt.Println("changed s is ", s)
}
