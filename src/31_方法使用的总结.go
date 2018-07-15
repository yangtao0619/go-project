// 31_方法使用的总结
package main

import (
	"fmt"
)

//func (temp long) testFunc() {
//	fmt.Println("temp is ", temp)
//}
type Person struct {
	name string
	age  int
}

func (temp *Person) testFunc() {
	fmt.Println("2 temp is ", temp)
}

func main() {
	fmt.Println("Hello World!")
	var person *Person = &Person{"Mike", 23}

	person.testFunc()
}
