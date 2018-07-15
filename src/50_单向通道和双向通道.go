// 50_单向通道和双向通道
package main

//"fmt"

func main() {
	var shuangChannel = make(chan int)

	var danChannel chan<- int = shuangChannel

	danChannel <- 222
}
