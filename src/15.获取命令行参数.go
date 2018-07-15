// 15.获取命令行参数
package main

import (
	"fmt"
	"os"
)

func main() {
	//	获取参数切片
	list := os.Args
	fmt.Println("list 长度是 ", len(list))
	//对切片进行遍历
	for i, data := range list {
		fmt.Printf("list[%d] is %s\n", i, data)
	}

}
