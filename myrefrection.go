package main

import (
	"fmt"
	"reflect"
)

type person struct {
	sex string
	age int
}

func (p *person) talk() {
	fmt.Println("i can talk with you")
}

func main() {
	p1 := person{"male", 25}
	Getinfo(p1)
}

func Getinfo(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())

	v := reflect.ValueOf(o)
	fmt.Println("Fields:")

	for i := 0; i < v.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}
}
