// 43_从命令行接收参数拷贝文件
package main

import (
	"fmt"
	"io"
	"os"
)

func copyFile(srcFilePath string, destFilePath string) {
	//首先要打开一个文件
	sFile, srcErr := os.Open(srcFilePath)
	if srcErr != nil {
		fmt.Println("无法正常打开源文件,err is ", srcErr)
		return
	}

	buffer := make([]byte, 4*1024)

	//创建一个文件
	dFile, dErr := os.Create(destFilePath)
	if dErr != nil {
		fmt.Println("文件创建失败,err is ", dErr)
		return
	}

	for {
		_, rErr := sFile.Read(buffer)

		if rErr != nil && rErr != io.EOF {
			fmt.Println("文件读取出错,err is ", rErr)
			return
		}

		if rErr == io.EOF {
			fmt.Println("文件读取结束")
			break
		}

		dFile.Write(buffer[:n])

	}

	defer sFile.Close()
	defer dFile.Close()
}

func main() {
	//从命令行获取参数,得到源文件和目标文件
	list := os.Args
	for i, data := range list {
		fmt.Printf("list[%d] is %s\n", i, data)
	}

	srcPath := list[1]
	destPath := list[2]
	fmt.Println("srcPath is ", srcPath, " destpath is ", destPath)

	copyFile(srcPath, destPath)
}
