package main

import (
	"bufio"
	"fmt"
	"go-src/functional/fib"
	"os"
)

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	//panic("error eccurred")
	return
	fmt.Println(4)
}

func tryDeferMany()  {
	for i:=0;i<100 ;i++  {
		defer fmt.Println(i)
		if i==30 {
			panic("too many nums")
		}
	}
}

func WriteFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := fib.Fibnacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func WriteFileToError(filename string)  {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		//panic(err)
		//fmt.Println("file exits")
		fmt.Println("Error: ",err.Error())
		return
	}else {
		fmt.Println("open file seccessful ")
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := fib.Fibnacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	//tryDefer()
	//WriteFile("adc.txt")
	//tryDeferMany()
	WriteFileToError("abc.txt")
}
