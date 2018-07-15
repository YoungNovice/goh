package main 

import (
	 "fmt"
)

func fb(n int) int {
	if	n == 1 || n == 2 {
		return 1
	}
	var f int =  fb(n-1) + fb(n-2)
	fmt.Println("f = ", f)
	return f
}

func main() {
	var a int = 10
	fb(a) 
}