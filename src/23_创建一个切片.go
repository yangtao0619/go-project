// 23_创建一个切片
package main

import (
	"fmt"
)

func testSlice() {
	arr := []int{1, 2, 3, 8, 9, 10, 12, 23, 14, 25}
	slice1 := arr[2:5]
	slice1[2] = 777
	slice2 := slice1[2:7]
	slice2[4] = 900
	fmt.Println(arr, "\n", slice1, "\n", slice2)
}

func testSlice_01() {
	arr := [5]int{1, 2, 3, 9, 10}
	slice := arr[1:3:5]
	for i, data := range slice {
		fmt.Printf("slice[%d] is %d\n", i, data)
	}
	fmt.Printf("len is %d ,cap is %d", len(slice), cap(slice))
}

func main() {
	testSlice()
}
