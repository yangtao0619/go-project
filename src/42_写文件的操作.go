// 42_写文件的操作
package main

import (
	"fmt"
	"os"
)

func writeFile(path string) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("文件创建失败,err = ", err)
		return
	}

	//文件创建成功之后需要向里面写入内容
	for i := 0; i < 10; i++ {
		file.WriteString("haha\n")
	}
	defer file.Close()
}

func main() {
	writeFile("./demo.txt")
}
