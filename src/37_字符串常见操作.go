// 37_字符串常见操作
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("hello world", "hello"))

	//组合
	slice := []string{"haha", "heihei", "gugu", "lala"}
	slice2 := strings.Join(slice, "|")
	fmt.Println("slice2 is ", slice2)

	//Trim函数和Feilds函数的区别 Trim去掉两头的空格,或者特定字符串,Fields去除空格,
	//并返回一个切片
	str := "    are    you   ok  ?"
	fmt.Println(strings.Trim(str, " "))

	fmt.Println(strings.Fields(str))

	//bool转换成string
	str2, err := strconv.ParseBool("false")
	fmt.Println(str2, err)
	//整型转换成字符串

	fmt.Println(strconv.Itoa(32))
	//string转换成int

	fmt.Println(strconv.Atoi("32"))
}
