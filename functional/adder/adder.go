package main

import "fmt"

//闭包
func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

//稍微正统的函数式编程
type Iadder func(int) (int, Iadder)

func adder2(base int) (Iadder) {
	return func(v int) (int, Iadder) {
		return (base + v), adder2(base + v )
	}
}

func main() {
	//a := adder()
	a := adder2(0)
	for i := 0; i < 10; i++ {
		//fmt.Println(a(i))
		var s int
		s ,a = a(i)
		fmt.Println(s)
	}
}
