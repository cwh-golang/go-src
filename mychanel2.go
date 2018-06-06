package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(runtime.NumCPU())
	c := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go foo(c, i)
	}
	for i := 0; i < 10; i++ {
		<-c
	}
}

func foo(c chan bool, index int) {
	s := 0
	for i := 0; i < 100000000; i++ {
		s += i
	}
	fmt.Println(index, s)
	// if index == 9 {
	// 	c <- true
	// }
	//  c <- true
	// if index == 9 {
	// 	close(c)
	// }
	c <- true
}
