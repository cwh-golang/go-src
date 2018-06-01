package main

//用括号的方式一次性引入/定义多个
import (
	// io "fmt"
	. "fmt"
	"strconv"
)

var (
	v_int1 = 1
	v_int2 = 2
	v_int3 = 3
)

const (
	c_const1 = 3.14
	c_const2 = 1.732
	c_const3 = 100
	c_const4 = "Hello world!"
)

type (
	newtype1 int
	newtype2 int
	newtype3 int
	newtype4 int
)

const (
	a_const = iota
	b_const = iota
	c_const = iota
	d_const = iota
	e_const = iota
	f_const = iota
)

const (
	_          = iota
	KB float64 = (1 << (iota * 10))
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
LABLE:
	for i := 0; i < 10; i++ {
		for {

			Println(i)
			continue LABLE
		}
	}
	Println("OK")
	Println("我是中国人！")
	b_int := 10
	switch {
	case b_int >= 1:
		Println("b_int>=1")
		fallthrough
	case b_int >= 2:
		Println("b_int>=2")
		fallthrough
	default:
		Println("b_int>=3")
	}
	a_int := 10
	var p *int = &a_int
	a_int++
	Println((*p))
	Println(p)
	Println(&a_int)
	Println(*p)
	Println(&p)
	Println(1 > 5)
	Println("5.0 / 7.8 = ", 5.0/7.8)
	Println("a" + "b")
	Println(!true)
	Println(true && false)
	Println(YB)
	Println(ZB)
	Println(EB)
	Println(PB)
	Println(TB)
	Println(GB)
	Println(MB)
	Println(KB)
	Println(1 << (d_const * 10))
	Println(1 << (c_const * 10))
	Println(1 << (b_const * 10))
	Println(1 << (a_const * 10))
	Println(6 &^ 11)
	Println(6 ^ 11)
	Println(6 & 11)
	Println(6 | 11)
	Println(1 << 10 << 10 << 10 >> 20)
	// fmt.Println("Hello,World!你好，世界！")
	// io.Println("Hello.world!你好，中国！")
	Println(a_const)
	Println(b_const)
	Println(c_const)

	k := len("a_const")
	Println(k)

	var a int = 66
	b := string(a)
	Println(b)
	c := strconv.Itoa(a)
	Println(c)
	d, _ := strconv.Atoi(c)
	Println(d)
	Println("Hello world!你好中国！")
	Println(3)
	Println(v_int1)
	Println(v_int2)
	Println(v_int3)

	Println(c_const1)
	Println(c_const2)
	Println(c_const3)
	Println(c_const4)
}
