// 56_测试请求报文
package main

import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("net listen err is ", err)
		return
	}

	conn, err1 := ln.Accept()
	if err1 != nil {
		fmt.Println("ln accept err is ", err1)
		return
	}

	buffer := make([]byte, 2048)
	n, err2 := conn.Read(buffer)
	if err2 != nil {
		fmt.Println("conn read err is ", err2)
		return
	}

	fmt.Println("data is ", string(buffer[:n]))
}
