// range的使用
package main

import (
	"fmt"
)

func main() {
	str := "hellogo"
	for i, data := range str {
		fmt.Printf("第 %d 个字符是 %c\n", i, data)
	}
}
