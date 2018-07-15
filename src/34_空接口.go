// 34_空接口
package main

import (
	"fmt"
)

//空接口函数
func testEmptyInterface(args ...interface{}) {

}

type Student struct {
	name  string
	score int
}

func main() {
	slice := make([]interface{}, 3)
	slice[0] = 1
	slice[1] = "haha"
	slice[2] = Student{"Mike", 90}

	for i, data := range slice {
		switch value := data.(type) { //得到对应的切片元素的类型
		case int:
			fmt.Printf("slice[%d] i is int,type is %v\n", i, value)
			break
		case string:
			fmt.Printf("slice[%d] i is string,type is %v\n", i, value)
			break
		case Student:
			fmt.Printf("slice[%d] i is Student,type is %v\n", i, value)
			break
		}
	}
}
