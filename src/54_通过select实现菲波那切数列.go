// 54_通过select实现菲波那切数列
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)    //数字通信
	quit := make(chan bool) //确认程序是否需要退出

	go func() {
		for i := 0; i < 20; i++ {
			num := <-ch
			fmt.Println("num is ", num)
		}
		//循环结束,程序终止
		quit <- true
	}()

	fibonacci(ch, quit)
}

func fibonacci(channel1 chan<- int, channel2 <-chan bool) {
	x, y := 1, 1
	for {
		select {
		case channel1 <- x:
			x, y = y, x+y
		case flag := <-channel2:
			fmt.Println("flag is ", flag)
			return
		}

	}
}
