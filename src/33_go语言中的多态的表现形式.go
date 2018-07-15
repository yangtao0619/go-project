// 33_go语言中的多态的表现形式
package main

import (
	"fmt"
)

type Humaner interface {
	sayHi()
}

type Student struct {
}

func (student *Student) sayHi() {
	fmt.Println("Student say hi")
}

type Person struct {
}

func (person *Person) sayHi() {
	fmt.Println("Person say hi")
}

type Teacher struct {
}

func (teacher *Teacher) sayHi() {
	fmt.Println("Teacher say hi")
}

func whoSayHi(i Humaner) {
	i.sayHi()
}

func main() {
	var teacher *Teacher = &Teacher{}
	whoSayHi(teacher)

	var person *Person = &Person{}
	whoSayHi(person)

	var student *Student = &Student{}
	whoSayHi(student)
}
