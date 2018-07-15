// 55_os.stat的使用
package main

import (
	"fmt"
	"os"
)

func main() {
	list := os.Args
	if 2 != len(list) {
		fmt.Println("输入错误")
		return
	}
	fileInfo, err := os.Stat(list[1])
	if err != nil {
		fmt.Println("获取文件属性失败,err is ", err)
		return
	}

	fmt.Println("file name is ", fileInfo.Name())
	fmt.Println("file size is ", fileInfo.Size())
}
