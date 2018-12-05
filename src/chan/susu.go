package main 

import (
	"fmt"
)

func main() {
	// 生成自然数的
	intChan := make(chan int, 1000)
	// 存放结果的
	primeChan := make(chan int, 2000)
	// exitChan
	exitChan := make(chan bool, 4)
	
	// 开启一个goroutine向intChan中放数据
	go putNumber(intChan)
}

func putNumber(intChan chan int) {
	for i := 1; i < 8000; i++ {
		intChan<-i
	}
	// 在添加元素完毕后关闭chan 因为关闭的chan也是可以读的 只不过读到最后会有返回值
	close(intChan)
}
