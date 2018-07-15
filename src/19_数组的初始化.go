// 19_数组的初始化
package main

import (
	"fmt"
)

func main() {
	arr := [4]int{2: 12, 3: 22}
	fmt.Println("arr is ", arr)

	//二维数组的初始化
	secondArr := [2][3]int{1: {4, 5, 6}}
	fmt.Println("secondArr is ", secondArr)
}
