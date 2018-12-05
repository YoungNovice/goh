package main 

import (
	"fmt"
)

// close chan 关闭之后不能写了， 但是还可以读了
func main() {
	c := make(chan int, 3)
	c <- 3
	c <- 3
	close(c)
//	c <- 3
	fmt.Println("关闭管道")
	xx := <- c
	fmt.Println("xxx=", xx)

	intchan2 := make(chan int, 200)
	for i := 0; i < 100; i++ {
		intchan2 <- i*2
	}

	// 这样取值是有问题的 len是一个变化的值
//	for i := 0; i<len(intchan2); i++ {
//		fmt.Println("i= ", i, " intchan2 <- = ", <-intchan2)
//	}
	len := len(intchan2)
	for i := 0; i<len; i++ {
		fmt.Println("i= ", i, " intchan2 <- = ", <-intchan2)
	}
	
}
