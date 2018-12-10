package main

import (
	"fmt"
)



// iota是怎么算出来的 其实一行就是一个递增 不管跟iota有没有关系
func main() {
	const (
		a = iota
		b 
		_
		c = 0
		j = -1
		k = -2 
		d = iota
		e, f = iota, iota
	)

	fmt.Println(a, b, c, d, e, f)

}
