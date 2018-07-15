// 24_创建一个猜数字的小程序
package main

import (
	"fmt"
	"math/rand"
	"time"
)

//此程序用来创建一个随机的四位数
func createRandNum() (result int) {
	var num int
	rand.Seed(time.Now().UnixNano())

	//创建一个四位数
	for {
		num = rand.Intn(10000)
		//判断一下,如果随机生成的数大于等于1000的话就中断循环
		if num >= 1000 {
			fmt.Println("num is ", num)
			break
		}
	}
	return num
}

//创建一个函数,将一个四位数分割成四个数字
func devideNum(num int) {
	//输入的num必须是一个四位数,需要先取出每一位
	qian := num / 1000
	bai := (num / 100) % 10
	shi := (num / 10) % 10
	ge := num % 10

	fmt.Println(qian, bai, shi, ge)
}

func main() {
	num := createRandNum()
	devideNum(num)
}
