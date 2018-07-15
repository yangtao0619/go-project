// Client
package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

/*
	该函数用来将本地的文件写入服务器中
*/
func uploadFile(path string, conn net.Conn) {
	srcFile, err := os.Open(path)
	if err != nil {
		fmt.Println("open src file err, ", err)
		return
	}

	defer srcFile.Close()

	buffer := make([]byte, 2048)

	//循环读取文件,读了多少,就往服务器写多少
	for {
		n, err2 := srcFile.Read(buffer)
		if err2 != nil {
			fmt.Println("src file read err", err2)
			if err2 == io.EOF {
				break
			} else {
				continue
			}
		}
		conn.Write(buffer[:n])
	}

}

func main() {
	fmt.Println("请输入要传输的文件:")
	var path string
	fmt.Scan(&path)

	//get file name
	info, errStat := os.Stat(path)
	if errStat != nil {
		fmt.Println("err stat is ", errStat)
		return
	}

	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		fmt.Println("net dial err is ", err)
		return
	}

	defer conn.Close()

	n, err2 := conn.Write([]byte(info.Name()))
	if err2 != nil {
		fmt.Println("conn write err is ", err2)
		return
	}

	buffer := make([]byte, 2048)
	n, err3 := conn.Read(buffer)
	if err3 != nil {
		fmt.Println("conn read err is ", err3)
		return
	}
	fmt.Println("msg from server is ", string(buffer[:n]))
	if "ok" == string(buffer[:n]) {
		//如果服务器给与响应了,说明可以向其传输文件了
		uploadFile(path, conn)
	}

}
