package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("exit!!")
	os.Exit(3)
}
