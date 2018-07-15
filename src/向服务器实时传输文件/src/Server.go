// Server
package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	ln, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("net listen err is ", err)
		return
	}

	defer ln.Close()

	conn, err2 := ln.Accept()
	if err2 != nil {
		fmt.Println("ln accept err is ", err2)
		return
	}

	defer conn.Close()

	buffer := make([]byte, 2048)

	n, err3 := conn.Read(buffer)
	if err3 != nil {
		fmt.Println("conn read err is ", err3)
		return
	}
	//读取发送过来的文件名,然后返回一个ok
	fileName := string(buffer[:n])
	fmt.Println("将要接收的文件名是:", fileName)

	conn.Write([]byte("ok"))

	//然后继续接收输入
	downloadFile(fileName, conn)
}

func downloadFile(name string, conn net.Conn) {
	destFile, err := os.Create(name)
	if err != nil {
		fmt.Println("os create file err,", err)
		return
	}
	defer destFile.Close()

	buffer := make([]byte, 2048)
	for {
		n, err2 := conn.Read(buffer)
		if err2 != nil {
			if err2 == io.EOF {
				fmt.Println("文件接收完毕")
			} else {
				fmt.Println("conn.Read err = ", err2)
			}
			return
		}

		if n == 0 {
			fmt.Println("n == 0 文件接收完毕")
			break
		}

		destFile.Write(buffer[:n])
	}
}
