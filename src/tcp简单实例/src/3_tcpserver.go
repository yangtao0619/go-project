// 3_tcpserver
//需求是做一个简单的并发服务器,服务器接收客户端发来的字母组成的字符串并转换成大写返回
package main

import (
	"fmt"
	"net"
	"strings"
)

func ChangeWord(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 2048)
	for {
		n, err2 := conn.Read(buffer)
		if err2 != nil {
			fmt.Println("读取数据失败,err is ", err2)
		}
		if "exit" == string(buffer[:n-2]) {
			break
		}

		fmt.Println("addr is ", conn.RemoteAddr().String(), "data is ", string(buffer[:n]))

		conn.Write([]byte(strings.ToUpper(string(buffer[:n]))))
	}

}

func main() {
	ln, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("连接失败,err is ", err)
		return
	}

	for {
		conn, err1 := ln.Accept()
		if err1 != nil {
			fmt.Println("接收数据失败,err1 is ", err1)
			return
		}
		//		fmt.Println("addr ", ln.Addr(), " connect success!")
		go ChangeWord(conn)
	}

}
