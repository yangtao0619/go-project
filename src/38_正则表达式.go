// 38_正则表达式
package main

import (
	"fmt"
	"regexp"
)

func testRegex() {
	str := "abc bac acc avc anc aic aoc add bsc a7c a0c"
	//制定规则
	regex := regexp.MustCompile("a[0-9]c")

	if regex == nil {
		//说明解释失败
		fmt.Println("解释失败")
		return
	}

	//匹配规则
	result := regex.FindAllStringSubmatch(str, -1)

	fmt.Println("result = ", result)
}

func testRegex2() {
	str := "1.23 2.342 4.231 0.23 jsdh 8.2323 hjdso"

	regex := regexp.MustCompile("\\d+\\.\\d+")
	if regex == nil {
		//说明解释失败
		fmt.Println("解释失败")
		return
	}

	//匹配规则
	result := regex.FindAllStringSubmatch(str, -1)

	fmt.Println("result = ", result)
}

func main() {
	testRegex2()
}
