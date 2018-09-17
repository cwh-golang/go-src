package main

import (
	"fmt"
)

func intSeq() func() (int, int) {
	i := 1
	return func() (int, int) {
		j := 7
		i += 1
		j = i + j
		return i, j
	}
}

func main() {
	func1 := intSeq()

	fmt.Println(func1()) //the func1 is not released!
	fmt.Println(func1())
	fmt.Println(func1())
	// delete func1

	// func2 := intSeq()
	// // func2 := (func() int)&func1

	// fmt.Println(func2())
	// fmt.Println(func2) //the same func pointer
	// fmt.Println(func1)

}
