package main 

import (
	"fmt"
	"strconv"
)

// 这个匿名函数返回的值是直接 n + x 就有点反常
// 如果分两行写结果就不一样了 为什么

// AddUpper 的返回值是一个函数
// 但是这个匿名函数引用到了函数外的n 
// 因此这个匿名函数和这个n就形成了一个整体 这个整体叫做闭包
// 大家可以这样理解；闭包是类 函数是操作 n是字段
// 函数和它使用的到n构成闭包 
func AddUpper() func (int) int {
	var n int = 10
	var str string = "hello"
	return func (x int) int {
		n += x
		str += strconv.Itoa(x)
		fmt.Println("str = ", str)
		return n
	}
}

func main() {
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))

	// 每次AddUpper调用n 就初始化一次 
	// 匿名函数对n 有影响
	// 但是f 和 j 函数是属于两个闭包 他们各自累加各自的n
	j := AddUpper() 
	fmt.Println(j(10))
	fmt.Println(j(10))
	fmt.Println(j(10))

}