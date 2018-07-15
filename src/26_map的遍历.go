// 26_map的遍历
package main

import (
	"fmt"
)

func main() {
	//定义一个map
	var students map[int]string = map[int]string{1: "mike", 2: "jack", 3: "john"}
	for key, value := range students {
		fmt.Printf("key is %d,value is %s\n", key, value)
	}

	//删除某一元素
	delete(students, 2)
	fmt.Println("students is ", students)
}
