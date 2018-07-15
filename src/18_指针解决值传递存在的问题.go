// 18_指针解决值传递存在的问题
package main

import (
	"fmt"
)

func swap(p1, p2 *int) {
	*p1, *p2 = *p2, *p1
}

func main() {
	a := 10
	b := 22
	fmt.Println("a is ", a, " b is ", b)
	swap(&a, &b)
	fmt.Println("a is ", a, " b is ", b)
}
