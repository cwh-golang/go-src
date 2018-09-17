package main

import (
	"fmt"
	"sort"
)

func A() (a, b, c int) {
	a, b, c = 1, 2, 3
	return
}

func B(b ...int) {
	for i, _ := range b {
		b[i] = i + 3
	}
	fmt.Println(b)
}

func main() {
	a, b := 1, 2
	B(a, b)
	fmt.Println(a, b)

	fmt.Println(A())

	m2 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d"}
	m3 := make(map[string]int)
	s2 := make([]int, len(m2))
	s3 := make([]string, len(m2))
	j := 0
	for k, v := range m2 {
		s2[j] = k
		s3[j] = v
		j++
	}
	fmt.Println(s2)
	fmt.Println(s3)
	j = 0
	for i := 0; i < len(m2); i++ {
		m3[s3[i]] = s2[i]
	}
	fmt.Println(m3)
	fmt.Println(m3["b"])
	m1 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d"}
	fmt.Println(m1)
	s1 := make([]int, len(m1))
	i := 0
	for k, _ := range m1 {
		s1[i] = k
		i++
	}
	sort.Ints(s1)
	fmt.Println(s1)
	for i, _ := range s1 {
		fmt.Println(m1[s1[i]])
	}

	m := make(map[int]string)
	m[2] = "Good"
	for i, v := range m {
		fmt.Println(i)
		fmt.Println(v)
	}
	fmt.Println(m)
	m[1] = "Hello"
	m[1] = "world"
	//delete(m, 1)
	fmt.Println(m)
	fmt.Println("Hello map")
}
