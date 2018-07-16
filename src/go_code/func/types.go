package main 

import (
	"fmt"
)

type myint int

type myFunc func (int) int

func test(n int) int{
	fmt.Println("test")
	return 0
}
func (my *myint) Increase(x int) {
	(*my) += myint(x)
}

func (my myFunc) xxx(x int) {
	 ret := my(x + 1)
	 fmt.Println("xxx函数调用 ret= ", ret)
}

func main() {
	var myFunc 

	var a myint = 1
	a.Increase(100)
	a.Increase(50)
	fmt.Println(a)	
}