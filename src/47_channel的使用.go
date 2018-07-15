// 47_channel的使用
package main

import (
	"fmt"
	"time"
)

var channel = make(chan int)

//交替打印两个字符数组
func printer(text string) {
	for _, data := range text {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}

func thread1() {
	printer("hello")
	channel <- 666

}

func thread2() {
	<-channel
	printer("world")
}

func main() {
	//	go thread1()
	//	go thread2()
	//	for {
	//	}

	//创建管道
	var chann = make(chan string)

	go func() {
		defer fmt.Println("子协程结束")
		for i := 0; i < 8; i++ {
			fmt.Printf("%d \n", i)
			time.Sleep(time.Second)
			if i == 3 {
				chann <- "hello moto"
				break
			}
		}

	}()
	str := <-chann
	defer fmt.Println("主协程结束", str)
}
