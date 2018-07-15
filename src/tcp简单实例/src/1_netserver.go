// 1_testTcp
package main

import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("err is ", err)
		return
	}

	//阻塞等待用户输入
	conn, err1 := ln.Accept()
	if err1 != nil {
		fmt.Println("err1 is ", err1)
		return
	}
	buf := make([]byte, 1024)
	n, err2 := conn.Read(buf)
	if err2 != nil {
		fmt.Println("err2 is ", err2)
		return
	}

	defer conn.Close()
	fmt.Println("data is ", string(buf[:n]))
}
