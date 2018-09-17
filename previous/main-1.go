package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

//import "math/cmplx"

func main() {
	fmt.Println("Hello World,你好，世界")

	//resp, err := http.Get("http://www.baidu.com")
	//		resp, err := http.Get("http://example.com")
	//		if err != nil {
	//			fmt.Println("网络获取错误！")
	//		}
	//		defer resp.Body.Close()
	//		io.Copy(os.Stdout, resp.Body)

	//	resp, err := http.Post("http://example.com/upload", "image/jpeg", &imageDataBuff)
	//	if err != nil {
	//		fmt.Println("网络获取错误！")
	//		return
	//	}

	//	if resp.StatusCode != http.StatusOK {
	//		fmt.Println("error！")
	//		return
	//	}

	resp, err := http.Head("http://example.com")
	if err != nil {
		fmt.Println("网络获取错误！")
	}
	//defer resp.Close
	io.Copy(os.Stdout, resp.Body)
	http.Client
	fmt.Println()
	fmt.Println("我是一个粉刷匠，粉刷本领强！")
}
