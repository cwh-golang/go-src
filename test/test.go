package main

import (
	"fmt"
	"regexp"
)

func main() {
	//text := `<li><a href='/read/1382563432/'Hello 世界！123 Go.`
	text := `<li><a href='/read/1382563432/' target='_blank'><img src='http://www.jgxedu.cn/uploads/image/2015/05/29/20150529083755.jpg' alt='义兴乡政府到义兴小学检查指导工作' width='310' height='207'/></a>`
	// 查找行首以 H 开头，以空白结尾的字符串（非贪婪模式）
	// 查找以 hello 开头（忽略大小写），以 Go 结尾的字符串
	//reg := regexp.MustCompile(`(?i:^hello).*Go`)
	//reg := regexp.MustCompile(` ^http://([\w-]+\.)+[\w-]+(/[\w-./?%&=]*)?$`)
	reg := regexp.MustCompile(`[a-zA-z]+://[^\s]*(jpg|png|gif)`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["Hello 世界！123 Go"]
}
