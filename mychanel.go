package main

import (
	"fmt"
)

func main() {
	c := make(chan bool)
	go func() {
		fmt.Println("Hello ")
		c <- false
		c <- true
		close(c)
	}()
	// <-c
	for v := range c {
		fmt.Println(v)
	}
}
