// 5_switch语句
package main

import (
	"fmt"
)

func main() {
	num := 10
	switch num {
	case 1:
		fmt.Println("这是1楼")
		break
	case 10:
		fmt.Println("这是10楼")
		break
	default:
		break
	}

	score := 86
	switch {
	case score > 85:
		fmt.Println("优秀")
		break
	case score > 60:
		fmt.Println("及格")
		break
	}
}
