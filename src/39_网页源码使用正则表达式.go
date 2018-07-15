// 39_网页源码使用正则表达式
package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	str := `
	<p>简书居然没有官方 Markdown 教程，我来</p>
	<p>原来官方是有的。。。</p>
	<p>不过我这个更简。
	而且还有独门秘籍。</p>
	<p>首先，  “Markdown 其实很简 单。在简书上学习 Markdown 最方便。”</p>
	<p>为了获得上面的,在 Mark down 编辑器里输入：</p>
	<p>后最好加个空格。除此之外，还有 5 级标题，依次有不同的字体大小，即</p>
	`

	//指定一个正则表达式的规则
	regex := regexp.MustCompile(`<p>(?s:(.*?))</p>`)

	//使用正则规则来匹配字符串
	result := regex.FindAllStringSubmatch(str, -1)

	//将切片中的数据进行过滤
	for _, data := range result {
		fmt.Println("data is ", strings.Replace(data[1], " ", "", -1))
	}

}
