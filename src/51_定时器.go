// 51_定时器
package main

import (
	"fmt"
	"time"
)

func main01() {
	//timer创建一个定时器,在规定的时间后想通道里面写入当前的时间
	timer := time.NewTimer(2 * time.Second)

	fmt.Println("time now is", time.Now())

	t := <-timer.C

	fmt.Println("time is ", t)
}

func main() {
	//定时2s,2s后产生一个事件,向通道里面写内容
	<-time.After(2 * time.Second)
	fmt.Println("时间到")
}
