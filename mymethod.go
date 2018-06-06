package main

import (
	"fmt"
)

func main() {
	a1 := A{}
	b1 := B{}
	a1.A_print()
	b1.B_print()
	var c1 C
	c1 = 0
	c1.Increase()
	fmt.Println(c1)
	c1++
	fmt.Println(c1)
}

type A struct {
	Name string
}

type B struct {
	Name string
}

type C int

func (c *C) Increase() {
	*c += C(100)
}

func (a *A) A_print() {
	fmt.Println("A")
}

func (b *B) B_print() {
	fmt.Println("B")
}
