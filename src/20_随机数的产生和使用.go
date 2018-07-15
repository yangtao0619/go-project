// 20_随机数的产生和使用
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//产生随机数需要用到一个包,rand
	//在生成随机数之前需要先seed一下,如果每次传入seed的整型常量的值都一样的话,那么每次生成的
	//随机数都是一样的,因此这里建议使用当前的系统时间来作为种子
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		//newIntn(100)这里的100表示生成[0,100)之间的任意整数
		fmt.Println("rand is ", rand.Intn(100))
	}
}
