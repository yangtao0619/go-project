// 21_冒泡排序
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//随机产生一个数组
	var arr [10]int

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}

	//接下来进行排序
	var temp int
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				temp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}

	fmt.Println("arr is ", arr)

}
