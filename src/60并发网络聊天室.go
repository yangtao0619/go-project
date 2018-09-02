package main

import (
	"net"
	"fmt"
	"time"
	"strings"
)

/*
实现并发聊天服务器
功能简介:
1.多用户登录;
2.多用户管理,用户上线广播;
3.常用指令exit退出,listuser查看所有用户
4.超时退出
5.单用户通信,将客户端发送来的字符转换成大写并返回
 */

/*
子模块功能实现:
1.用户上线广播
全局map,全局message,处理单客户端连接的go程,处理message的go程,处理单客户channel的go程
 */

type ClientInfo struct {
	Name   string
	IpAddr string
	C      chan string
}

var message = make(chan string)

var onlineMap map[string]ClientInfo

func main() {
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("net listen err :", err)
		return
	}
	go ManageMessage()
	for {
		conn, err := listener.Accept()
		fmt.Println("There is a client log in")
		if err != nil {
			fmt.Println("listener Accept err :", err)
			return
		}
		go HandleClientConnect(conn)
	}
}

//管理全局消息,当全局消息来的时候,循环给每个客户端的管道赋值
func ManageMessage() {
	onlineMap = make(map[string]ClientInfo)
	for {
		msg := <-message
		for _, client := range onlineMap {
			client.C <- msg
		}
	}
}

func HandleClientConnect(conn net.Conn) {
	defer conn.Close()
	//这里需要将conn封装成client对象
	ipAddr := conn.RemoteAddr().String()
	client := ClientInfo{ipAddr, ipAddr, make(chan string)}
	//将client存储到map中
	onlineMap[ipAddr] = client
	//创建读取client.C的go程,先创建读,用于阻塞监听数据过来之后就可以直接读出来了,如果先写的话,可能会阻塞住
	go writeToClient(conn, client)
	//向message写入数据
	message <- "ip: [" + ipAddr + "] login!"
	//除了发送消息,还需要监听客户端发来的数据,进行消息的判断
	shouldQuit := make(chan bool)
	hasData := make(chan bool)
	go func() {
		buffer := make([]byte, 4096)
		for {
			n, err := conn.Read(buffer)
			if n == 0 {
				fmt.Println("检测到当前客户端退出,Client Name is", client.Name)
				return
			}

			if err != nil {
				fmt.Println("conn read err ", err)
				return
			}
			//读出来的数据解析一下
			str := string(buffer[:n])
			if str == "exit\n" {
				shouldQuit <- true
				break
			} else if str == "listuser\n" {
				//展示所有的用户消息写给当前的client
				fmt.Println(ipAddr + ": list all user")
				//循环的将当前的map中每个client的ip输出到message中
				for _, eachClient := range onlineMap {
					client.C <- "name::[" + eachClient.Name + "],ip::[" + eachClient.IpAddr + "]"
				}
			} else if len(str) >= 9 && strings.HasPrefix(str, "rename|") {
				//给当前的用户改名字
				str = strings.Replace(str, "\n", "", -1)
				split := strings.Split(str, "|")
				newName := split[1]
				client.Name = newName
				//将修改后的数据重新保存到map中
				onlineMap[ipAddr] = client
				client.C <- "rename to " + newName + " success!"
			} else {
				//将小写转成大写,写回给客户端
				str = strings.ToUpper(str)
				conn.Write([]byte(str))
			}
			hasData <- true
		}
	}()

	for {
		select {
		case <-shouldQuit:
			delete(onlineMap, ipAddr)
			//退出当前的客户端
			message <- "client " + client.Name + " exit!"
			//关闭用户自带的channel,结束writeToClient这个go程
			close(client.C)
			return
		case <-time.After(time.Second * 60):
			delete(onlineMap, ipAddr)
			//超时退出
			message <- "client " + client.Name + " time out!"
			//关闭用户自带的channel,结束writeToClient这个go程
			close(client.C)
			return
		case <-hasData:
			//啥也不做,用于重置计时器
		}
	}

}

/*
将数据写给客户端
 */
func writeToClient(conn net.Conn, client ClientInfo) {
	for msg := range client.C {
		conn.Write([]byte(msg + "\n"))
	}

	defer fmt.Println("write to client end")

}
