// 40_通过结构体生成json,同时对结构体进行二次编码
package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string    `json:"name"`
	Age   int       `json:"age",string`
	Sex   byte      `json:"sex",string`
	Score []float64 `json:"score"`
}

func main() {
	//声明一个结构体变量
	student := Student{"Mike", 23, 'M', []float64{89.9, 78.9, 90.5}}

	//通过结构体变量来完成json的输出
	buf, err := json.MarshalIndent(student, "", " ")

	if err != nil {
		fmt.Println("转换成json发生错误... ,", err)
		return
	}
	//转换成的是byte数组,需要再次转换成string
	fmt.Println("buf is ", string(buf))
}
