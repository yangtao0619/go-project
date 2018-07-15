// 29_匿名字段的初始化
package main

import (
	"fmt"
)

type Person struct {
	age  int
	sex  byte
	name string
}

type Student struct {
	Person
	id    int
	score int
}

func main() {
	//匿名字段的初始化
	var student Student = Student{Person: Person{age: 23}, id: 1}
	fmt.Printf("student is %+v \n", student)
}
