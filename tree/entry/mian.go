package main

import (
	"fmt"
	"go-src/queue"
)

func main() {
	q:=queue.Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	q.Push("hello world!")
	fmt.Println(q.Pop())
	arr:=make([]int,10,20)
	arr = append(arr, 1, )
	arr = append(arr, 2, )
	arr = append(arr, 2, )
	arr = append(arr, 2, )
	arr = append(arr, 2, )
	fmt.Println(arr)
	q.Push(arr)
	fmt.Println(q.Pop())
}
