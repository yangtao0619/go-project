package main

import "fmt"

/*
写下判断闰年的表达式,设待判断的年份变量为year.
闰年的判定(符合下面两个条件之一):
年份能够被400整除.(2000)
年份能够被4整除但不能被100整除.(2008)
让用户输入一个年份,如果是润年,则输出true,如果不是,则输出false.
2100/1600/1800/2009年是闰年吗？
*/

func main() {
	var year int
	fmt.Scanf("%d", &year)
	result := judgeLeapYear(year)
	if result {
		fmt.Printf("%d 年是闰年\n", year)
	} else {
		fmt.Printf("%d 年是平年\n", year)
	}

	arr := [4]int{2100,1600,1800,2009}
	for _,data := range arr{
		if judgeLeapYear(data){
			fmt.Printf("%d 年是闰年\n", data)
		}else{
			fmt.Printf("%d 年是平年\n", data)
		}
	}

}

func judgeLeapYear(year int) bool {
	//对year赋值之后可以进行判断
	if year%400 == 0 || ((year%4 == 0) && (year%100 != 0)) {
		return true
	} else {
		return false
	}
}
