package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(runtime.NumCPU())
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go foo(&wg, i)
	}
	wg.Wait()
}

func foo(wg *sync.WaitGroup, index int) {
	s := 0
	for i := 0; i < 100000000; i++ {
		s += i
	}
	fmt.Println(index, s)
	wg.Done()
}
