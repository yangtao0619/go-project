// 2_netclient
package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		fmt.Println("请求出错,err is ", err)
		return
	}

	conn.Write([]byte("hello go!"))
	defer conn.Close()
}
