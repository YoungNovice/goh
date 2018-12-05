package main

import (
	"fmt"
	"time"
)

// goroutine的painc会导致整个主线程的崩溃
// 在goroutine中使用recover处理painc问题
func main() {
	exitChan := make(chan int)
	go sayHello()
	go test()
	go doExit(exitChan)
	<-exitChan

}

func doExit(exitChan chan<- int) {
	time.Sleep(time.Second * 3)
	exitChan<-0
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("当前线程出现故障, 已经恢复处理", err)
		}
	}()
	var m map[int]string
	m[0] = "he"
}

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello,world", i)
	}
}
