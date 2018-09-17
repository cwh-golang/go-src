package fib


//为函数实现接口
type intGen func() int

//0 ,1, 1, 2, 3, 5, 8, 13...
//      a  b
//         a  b
//            a  b
func Fibnacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
