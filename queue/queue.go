package queue

//type Queue []int{}   //对[]int定义别名，来进行封装
type Queue []interface{}   //对[]interface定义别名，可以在支持任何类型

func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue)IsEmpty() bool {
	return len(*q) == 0
}