package main

import (
	"fmt"
)

type person struct {
	Nane string
	Age  int
}

//匿名结构
type student struct {
	Nane    string
	Age     int
	contact struct {
		Num, city string
	}
}

type teacher struct {
	string
	int
}

func A(per person) {
	per.Age = 15
	fmt.Println(per.Age)
}

func B(per *person) {
	per.Age = 15 //这里为什么不用*per.Age = 15
	fmt.Println(per.Age)
}

func main() {
	t1 := teacher{"Mark", 66}
	fmt.Println(t1)

	s1 := student{
		Nane: "tom",
		Age:  24,
	}
	s1.contact.Num = "1234"
	s1.contact.city = "chongqing"
	fmt.Println(s1)

	// b := &struct {
	b := struct {
		Name string
		Age  int
	}{
		Name: "lili",
		Age:  23,
	}
	fmt.Println(b)

	a := person{
		Nane: "joe",
		Age:  19,
	}
	// a.Nane = "joe"
	// a.Age = 19
	fmt.Println(a)
	// A(a)
	B(&a)
	fmt.Println(a)
}
