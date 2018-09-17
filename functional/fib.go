package main

import (
	"bufio"
	"fmt"
	"go-src/functionalctional/fib"
	"io"
	"strings"
)


//为函数实现接口
type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000{
		return 0,io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	//TODO : error if p is too small
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader)  {
	scaner:=bufio.NewScanner(reader)
	for scaner.Scan() {
		fmt.Println(scaner.Text())
	}
}

func main() {
	f := fib.Fibnacci()
	printFileContents(f)
}
