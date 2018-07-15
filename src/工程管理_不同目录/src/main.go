// mian
package main

import (
	"calc"
	"fmt"
)

func init() {
	fmt.Println("this is init of main")
}

func main() {
	result := calc.Add(10, 20)
	fmt.Println("result is ", result)
}
