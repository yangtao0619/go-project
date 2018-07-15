package main

import (
	"fmt"
	"net"
)

type Client struct {
	name string
	addr string
	C    chan string
}

var clientMap map[string]Client

var message = make(chan string)

func informLoginMsg() {
	//先给map分配空间
	clientMap = make(map[string]Client)
	//遍历map,给所有的管道赋值
	for {
		//这里死循环是为了一有消息的时候就遍历发送给所有的登录用户

		//在收取到新消息之前,这里会阻塞
		msg := <-message
		for _, value := range clientMap {
			value.C <- msg
		}
	}

}

func sendMsgToClient(client Client, conn net.Conn) {
	for msg := range client.C {
		conn.Write([]byte(msg))
	}
}

//处理客户端连接,需要向map中写入数据
func handleConn(conn net.Conn) {
	//连接上之后可以先向map中写入数据
	addr := conn.RemoteAddr().String()
	name := addr

	client := Client{name, addr, make(chan string)}
	clientMap[addr] = client
	message <- "ip " + addr + " login \n"

	//给当前的客户端发送消息
	go sendMsgToClient(client, conn)
	client.C <- addr + "|" + name + " i am here\n"

	//本地输出连接信息
	fmt.Println("ip ", addr, " | ", name, " login ")

	buffer := make([]byte, 2048)
	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("conn read err is ", err)
		return
	}

}

func main() {
	//当有用户连接上服务器之后,向所有的用户进行广播
	ln, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("net listen err is ", err)
		return
	}

	defer ln.Close()
	//当有客户端连接到服务器的时候,要广播给所有的用户
	go informLoginMsg()
	for {
		//循环读取接收到的数据
		conn, err2 := ln.Accept()
		if err2 != nil {
			fmt.Println("ln accept err is ", err2)
			return
		}

		defer conn.Close()

		//还需要新开一个协程,处理用户的连接
		go handleConn(conn)
	}

}
