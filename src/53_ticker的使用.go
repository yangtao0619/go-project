// 53_ticker的使用
package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(2 * time.Second)

	for i := 0; i < 4; i++ {
		<-ticker.C
		fmt.Println("i is ", i)

		if i == 3 {
			break
		}
	}

}
