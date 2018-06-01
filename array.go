package main

//用括号的方式一次性引入/定义多个
import (
	// io "fmt"
	"fmt"
	//"strconv"
)

func main() {
	a1 := [10]int{1, 2, 3, 4, 5}
	s4 := a1[:]
	//fmt.Printf("%p,%p\n", &a1, s4)
	s2 := make([]int, 3, 10)
	s2 = append(s2, 1, 2, 3)
	a_int := 4
	s2 = append(s2, a_int)
	// fmt.Printf("%p\n", s2)
	//	fmt.Println(len(s4), cap(s4))
	// fmt.Println(s4)
	fmt.Println(s2)
	copy(s2, s4)
	fmt.Println(s2)

	s3 := s2[3:]
	fmt.Printf("%p,%p\n", s2, s3)
	fmt.Println(len(s3), cap(s3))
	fmt.Println(s3)
	s2[1] = 2
	fmt.Println(s2)
	fmt.Println(len(s2), cap(s2))

	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var s1 []int
	fmt.Println(s1)
	s1 = a[5:]
	fmt.Println(s1)

	s1 = a[1:]
	fmt.Println(s1)

	arr := [...]int{3: 1}
	fmt.Println(arr)
	arr_2 := [3][2]int{
		{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(arr_2)
	fmt.Println(len(arr_2))
	p := new([3]int)
	fmt.Println(p)
}
