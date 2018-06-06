// muti_parameters.go

package main

//用括号的方式一次性引入/定义多个
import (
	// io "fmt"
	"fmt"
	//"strconv"
)

func sum(nums ...int) {
	fmt.Print(nums, " ")         //it's a slice
	nums = append(nums, 4, 5, 6) //this is the proof
	sum1 := 0
	for i := 0; i < len(nums); i++ {
		// fmt.Println(nums[i])
		sum1 += nums[i]
	}
	fmt.Println(sum1)
}

func main() {
	sum(1, 2, 3)
	sum(4, 5, 6, 7, 8)
}
