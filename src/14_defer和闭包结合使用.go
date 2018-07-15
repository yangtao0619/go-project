// 14_defer和闭包结合使用
package main

import (
	"fmt"
)

func main() {
	var a int
	var b int

	a = 1
	b = 1

	//这里面闭包捕获的是形参里面的a和b,第一次赋值的时候已经有了
	defer func(a int, b int) {
		fmt.Println("defer a + b is ", a+b)
	}(a, b)

	a = 12
	b = 23
	fmt.Println("a + b is ", a+b)

}
