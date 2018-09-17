// multi_result.go

package main

//用括号的方式一次性引入/定义多个
import (
	// io "fmt"
	"fmt"
	//"strconv"
)

func sum(a, b int) (res, err int) {
	if a == 0 || b == 0 {
		res = 0
		err = 1
	} else {
		res = a + b
		err = 0
	}
	return
}

func main() {
	fmt.Println(sum(1, 2))
	_, err := sum(3, 4)
	fmt.Println(err)
}
