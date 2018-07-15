package main 

import (
	"fmt"
)

func main() {
	a := 2
	b := 3
	// 2的补码 0000 0010
	// 3的补码 0000 0011
	// 2 & 3  0000 0010 2
	// 2|3    0000 0011 3 
	// 2^ 3   0000 0001 1
	// 2 &^ 3 1111 1101 FD 
	fmt.Println("2 & 3 ", a & b)
	fmt.Println("2 | 3", a | b)
	fmt.Println("2 ^ 3", a ^ b)
	fmt.Println("2 &^ 3", a &^ b)
}