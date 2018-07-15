// 32_方法的继承
package main

import (
	"fmt"
)

type Person struct {
	age  int
	name string
	sex  byte
}

func (p *Person) printPerson() {
	fmt.Printf("person is %+v\n", p)
}

type Student struct {
	Person
	id    int
	score int
}

func main() {
	//验证方法的继承
	var student *Student = &Student{Person{23, "Mike", 'm'}, 666, 87}
	student.printPerson()
}
