package main

import (
	"fmt"
	"io"
	"go-src/retrievertriever/mock"
	"go-src/retrievertriever/realre"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string,
		form map[string]string,
	)string
}

type RetrieverPoster interface {
	Retriever
	Poster
}

const url  = "http://www.imooc.com"

func post(p Poster)  string{
	return p.Post(url,
		map[string]string{
			"ccmouse":"imooc",
			"course":"golang"})
}

func downlord(r Retriever) string {
	return r.Get(url)
}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case realre.Reretriever:
		fmt.Println("userAgent:", v.UserAgent)
	case *mock.Retriever:
		fmt.Println("contents:", v.Contents)
	}
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents":"another fake imooc.com",
	})
	return s.Get(url)
}

func myread(r io.Reader)  {
	n, err := r.Read([]byte("hello wolrd!"))
	if err!=nil {
		fmt.Println("You have get a wrong parameter!")
	}
	fmt.Println(n)
}


func main() {
	var r Retriever
	mockRetriever := mock.Retriever{"This is fake imooc.com."}
	fmt.Println(&mockRetriever)
	r=&mockRetriever
	//fmt.Println(downlord(r))
	//fmt.Printf("%T %v\n", r, r) //%T 打印类型 %v打印内容
	r = realre.Reretriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	//fmt.Printf("%T %v\n", realR, realR)
	//fmt.Println(downlord(realR))

	inspect(r)
	inspect(r)

	//type assertion  类型断言
	//	reretriever := r.(mock.Retriever)
	//	fmt.Println(reretriever.Contents)

	if i, ok := r.(*mock.Retriever); ok {
		fmt.Println(i.Contents)
	} else {
		fmt.Println("not a realerRetriver")
	}

	rMock:=&mock.Retriever{
		"hello world",
	}

	fmt.Println(post(rMock))

	fmt.Println(session(&mockRetriever))

	myread(&mockRetriever)
}
