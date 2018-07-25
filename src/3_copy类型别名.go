// 3_copy类型别名
package main

import (
	"fmt"
)

type bigInt int64

func main() {
	var num bigInt = 77
	fmt.Println("num is ", num)
}
