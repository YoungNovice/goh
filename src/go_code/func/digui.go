package main 

import (
	"fmt"
)

func test(n int) {
	if n > 2 {
		n--
		test(n)
	} else {
		fmt.Println("n= ", n)
	}
}
// in if test(3)
// in if test(2)

// n = 2
// n = 2
// n = 3
func main() {
	test(4)
}