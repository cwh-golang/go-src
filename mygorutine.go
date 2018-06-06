package main

import (
	"fmt"
	"time"
)

func main() {
	go foo()
	time.Sleep(2 * time.Second)
}

func foo() {
	fmt.Println("go go go !!!")
}
