package main 

import (
	"fmt"
)

func f(n int) int {
	if n == 1 {
		return 3
	} else {
		return 	2 * (n-1) + 1
	}
}

func main() {
	fmt.Println(f(3))
}