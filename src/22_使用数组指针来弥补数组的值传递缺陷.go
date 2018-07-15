// 22_使用数组指针来弥补数组的值传递缺陷
package main

import (
	"fmt"
)

func change(arr [5]int) {
	arr[0] = 21
	fmt.Println("arr[0] is ", arr[0])
}

func changeP(arr *[5]int) {
	(*arr)[0] = 33
	fmt.Println("arr[0] is ", *arr)
}

func main() {
	var arr [5]int = [5]int{1, 2, 23, 12, 11}
	changeP(&arr)
	fmt.Println("arr[0] is ", arr[0])
}
