// 4_tcpclient

//需求描述,可以接收键盘输入和服务端返回数据的客户端
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		fmt.Println("net.Dial err is ", err)
		return
	}

	defer conn.Close()

	go func() {
		strBuffer := make([]byte, 2048)
		for {
			n, err3 := os.Stdin.Read(strBuffer)
			if err3 != nil {
				fmt.Println("os.Stdin.Read err is", err3)
				return
			}

			//将输入的数据写到服务器上面
			conn.Write(strBuffer[:n])
		}

	}()

	//取出服务器返回的数据
	buffer := make([]byte, 2048)

	for {
		n, err2 := conn.Read(buffer)
		if err2 != nil {
			fmt.Println("conn.Read err is ", err2)
			return
		}

		fmt.Println("data is ", string(buffer[:n]))
	}

}
