// 8函数测试
package main

import (
	"fmt"
)

//定义一个有参数没有返回值的函数
func testFunc(str string) {
	fmt.Println("str is ", str)
}

func testFunc2(args ...int) {
	fmt.Println("这里是不定参数", "元素个数总共有", len(args))
	//打印里面的每一个元素
	//	for i := 0; i < len(args); i++ {
	//		fmt.Printf("arg[%d] is %d\n", i, args[i])
	//	}

	//	for i, data := range args {
	//		fmt.Printf("arg[%d] is also %d\n", i, data)
	//	}

	testFunc3(args[2:]...)

	testFunc3(args[:2]...)
}

//不定参数的传递
func testFunc3(temp ...int) {
	for i, data := range temp {
		fmt.Printf("arg[%d] is also haha %d\n", i, data)
	}
}

func testMoreReturn(a int, b int, c int) (num int, sum int) {
	num = 3
	sum = a + b + c
	return
}

func main() {
	//	testFunc("hello go ,i am fine")
	//	testFunc2(2, 3, 2, 442, 123, 12)
	num, sum := testMoreReturn(2, 3, 4)
	fmt.Printf("num is %d ,sum is %d", num, sum)
}
