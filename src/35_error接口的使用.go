// 35_error接口的使用
package main

import (
	"errors"
	"fmt"
)

func testDev(a, b int) (result int, err error) {
	if a == 0 {
		err = errors.New("分母不能为0")
	} else {
		result = b / a
	}
	return
}

func main() {
	//	var err1 error = fmt.Errorf("%s ", "this is fmt error info")
	//	fmt.Println("err1 is ", err1)

	//	err2 := errors.New("this is errors info")
	//	fmt.Println("err2 is ", err2)
	result, err := testDev(1, 2)
	if err != nil {
		fmt.Println("err is ", err)
	} else {
		fmt.Println("result is ", result)
	}
}
