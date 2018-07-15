// 10_函数类型
package main

import (
	"fmt"
)

//定义函数类型,函数本身也可以当成一个类型来使用
type FuncType func(int, int) int

func add(a int, b int) (result int) {
	return a + b
}

func main() {
	//定义一个函数变量
	var testFunctype FuncType
	testFunctype = add
	result := testFunctype(3, 7)
	fmt.Println("result is ", result)

}
