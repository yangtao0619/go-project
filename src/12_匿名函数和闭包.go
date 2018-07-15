// 12_匿名函数和闭包
package main

import (
	"fmt"
)

func main() {
	a := 10
	str := "mike"
	//定义匿名函数,然后调用
	testFunc := func() {
		fmt.Println("a is ", a, " str is ", str)
	}
	testFunc()

	//定义匿名函数,同时调用
	func() {
		fmt.Printf("str is %s,haha\n", str)
	}()

	//带参数的匿名函数
	func(x int, y int) {
		fmt.Println("result is ", x+y)
	}(3, 4)

	//匿名函数,有参有返回值
	max, min := func(i int, j int) (max int, min int) {
		if i > j {
			return i, j
		} else {
			return j, i
		}
	}(80, 30)

	fmt.Printf("max is %d,min is %d", max, min)
}
