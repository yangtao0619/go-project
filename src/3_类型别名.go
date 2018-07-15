// 3_类型别名
package main

import (
	"fmt"
)

func main() {
	type bigInt int64 //给int64起一个别名叫bigInt
	var a bigInt = 21
	fmt.Printf("a type is %T", a)

}
