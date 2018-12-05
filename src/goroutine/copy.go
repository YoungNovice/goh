package main 

import (
	"fmt"
)


func writeData(intchan chan int) {
	for i := 1; i<= 50; i++ {
		intchan <- i
		fmt.Println("writedata i= ", i)
	}
}

func readData(intchan chan int, exitchan chan bool) {
	for i := 1; i<= 50; i++ {
		v := <-intchan
		fmt.Println("readdata v= ", v)
	}
	// 向退出管道写数据	
	exitchan <- true
}

// 这种chan需要自己搞个死循环来阻塞 太sb了 
// 而且如果不close的读会死锁的
func main() {
	intchan := make(chan int, 50)
	extichan := make(chan bool, 1)
	
	go writeData(intchan)
	go readData(intchan, extichan)
	<-extichan
}
