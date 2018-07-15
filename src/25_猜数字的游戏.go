// 25_猜数字的游戏
package main

import (
	"fmt"
	"math/rand"
	"time"
)

//需求描述,生成一个随机的四位数,从命令行输入以后和原来的数做对比

//生成四位随机数的函数
func createRandNum() (randNum int) {
	rand.Seed(time.Now().UnixNano())

	for {
		randNum = rand.Intn(10000)
		if randNum > 999 {
			break
		}
	}
	return
}

//从命令行接收一个四位数
func getNumFromCmd() (inputNum int) {
	for {
		fmt.Scan(&inputNum)
		//对输入的数字进行检测,如果是合法的四位数,才返回
		if inputNum > 999 && inputNum < 10000 {
			return
		} else {
			fmt.Println("输入的数字不合法,请重新输入")
		}
	}
}

//定义一个函数,将一个数字拆分成一个切片
func getSliceFromNum(num int, slice []int) {
	slice[0] = num / 1000
	slice[1] = num / 100 % 10
	slice[2] = num / 10 % 10
	slice[3] = num % 10

}

func main() {
	//生成一个随机的四位数
	randNum := createRandNum()
	fmt.Println("randNum is ", randNum)

	//拿输入的数和随机生成的数按位比较
	//创建一个切片
	slice_RandNum := make([]int, 4)
	getSliceFromNum(randNum, slice_RandNum)
	fmt.Println("slice_RandNum is ", slice_RandNum)

	for {
		//从命令行接收一个四位数
		inputNum := getNumFromCmd()
		fmt.Println("输入的数字是 ", inputNum)

		//将从命令行输入的数字切分成切片
		slice_InputNum := make([]int, 4)
		getSliceFromNum(inputNum, slice_InputNum)
		fmt.Println("slice_InputNum is ", slice_InputNum)
		var n int
		//两个切片按位进行比较
		for i, data := range slice_RandNum {
			if data > slice_InputNum[i] {
				fmt.Printf("第%d位输入的有点小\n", i+1)
			} else if data < slice_InputNum[i] {
				fmt.Printf("第%d位输入的有点大\n", i+1)
			} else {
				fmt.Printf("第%d位输入正确\n", i+1)
				n++
			}
		}
		if n == 4 {
			break
		}
	}

}
