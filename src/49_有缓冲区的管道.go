// 49_有缓冲区的管道,测试管道的关闭
package main

import (
	"fmt"
)

func main() {
	//创建一个有缓冲区的管道
	var bufferChannel = make(chan int, 3)

	go func() {
		for i := 0; i < 8; i++ {
			fmt.Println("子协程")
			bufferChannel <- i
		}
		close(bufferChannel)
	}()

	for {
		data, isok := <-bufferChannel
		if isok == true {
			//说明有数据
			fmt.Println("data is ", data)
		} else {
			break
		}
		//		fmt.Println("主协程")
	}

}
