package main 

import (
	"fmt"
)

func test1() int {
	return 1
}

func test2(x int) int {
	return 1 + x
}
func main() {
	// case 后面可以有多个表达式 ， 少用fullthrough
	// switch 和 case 里面都是表达式  就是可以比较的两个值
	// switch 后面还可以不带表达式 类似于 if else
	// switch 后面还可以声明一个变量 但是后面要加分号	
	var a = 0
	switch test1() {
	case test2(a), 5 :
		fmt.Println("aaaa") 
	default :
	fmt.Println("没有值") 
	}

}