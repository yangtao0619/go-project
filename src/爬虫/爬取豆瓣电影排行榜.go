package main

import (
	"fmt"
	"strconv"
	"net/http"
	"io"
	"regexp"
	"os"
	_ "github.com/modood/table"
	"github.com/modood/table"
)

/*
功能分析:
1.首先确定要爬取的页面的url之间的规律
https://movie.douban.com/top250?start=0&filter=
https://movie.douban.com/top250?start=25&filter=
https://movie.douban.com/top250?start=50&filter=   https://movie.douban.com/top250?start=(i-1)*25&filter=
 */

func main() {
	//提示用户输入起始页码和终止页码
	var start, end int
	fmt.Println("请输入起始页:")
	fmt.Scan(&start)
	fmt.Println("请输入终止页:")
	fmt.Scan(&end)
	DoWork(start, end)

}

type movie struct {
	Name       string
	Score      float64
	MarkNumber int
}

func DoWork(start int, end int) {
	channel := make(chan int)

	//开启go程循环的爬取数据
	for i := start; i <= end; i++ {
		url := "https://movie.douban.com/top250?start=" + strconv.Itoa((i-1)*25) + "&filter="
		go spiderPage(i, url, channel)
	}

	for i := start; i <= end; i++ {
		fmt.Printf("第%d页爬取完毕\n", <-channel)
	}
}
func spiderPage(index int, url string, channel chan<- int) {
	//得到数据之后需要对数据进行分析,会得到三个二维数组,把这几个数组每行下标为1的数据取出来存在另一个数组里面
	data := getDateByUrl(url)
	//电影名称
	//<img width="100" alt="肖申克的救赎" src="
	nameRegex := regexp.MustCompile(`<img width="100" alt="(?s:(.*?))" src="`)
	names := nameRegex.FindAllStringSubmatch(data, -1)
	//电影评分

	//<span class="rating_num" property="v:average">9.6</span>
	scoreRegex := regexp.MustCompile(`<span class="rating_num" property="v:average">(?s:(\d\.\d*?))</span>`)
	scores := scoreRegex.FindAllStringSubmatch(data, -1)
	//评价人数
	//<span>1131909人评价</span>
	markNumberRegex := regexp.MustCompile(`<span>(?s:(\d*?))+人评价</span>`)
	markNumbers := markNumberRegex.FindAllStringSubmatch(data, -1)
	//每一页的数据都需要处理存储到文件里面
	storeDataToFile(index, names, scores, markNumbers)
	channel <- index
}
func storeDataToFile(index int, names [][]string, scores [][]string, markNumbers [][]string) {
	//创建文件
	fileName := "第" + strconv.Itoa(index) + "页.txt"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os create err:", err)
		return
	}
	defer file.Close()
	//file.WriteString("电影名称" + "\t\t" + "电影评分" + "\t\t" + "评价人数\n")
	//开启循环
	movies := make([]movie, 0)
	for i := 0; i < len(names); i++ {
		name := names[i][1]
		score := scores[i][1]
		markNumber := markNumbers[i][1]

		scoreNew, err := strconv.ParseFloat(score, 64)
		if err != nil {
			fmt.Println("strconv.ParseFloat:", err)
			continue
		}
		markNumberNew, err := strconv.Atoi(markNumber)
		if err != nil {
			fmt.Println("strconv.Atoi:", err)
			continue
		}
		movies = append(movies, movie{name, scoreNew, markNumberNew})
		//file.WriteString(Name + "\t\t" + Score + "\t\t" + MarkNumber + "\n")
	}
	result := table.Table(movies)
	file.WriteString(result)
}

/*
使用url发起http get请求获取页面数据
 */
func getDateByUrl(url string) (result string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("http get err:", err)
	}

	defer response.Body.Close()
	buffer := make([]byte, 4096)
	for {
		n, err := response.Body.Read(buffer)
		if n == 0 {
			break
		}
		if err != nil && err != io.EOF {
			fmt.Println("response read err:", err)
		}
		result += string(buffer[:n])
	}
	return
}
