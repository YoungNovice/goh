package main 

import (
	"fmt"
	"time"
	"sync"
)

// 资源竞争问题怎么解决
// 线程间通讯的问题怎么解决

// 多个goroutine同时访问一份资源有并发问题
// go build -race xxx.go用来检测是否有并发问题
var myMap map[int]int64 = make(map[int]int64, 200)
var lock *sync.Mutex = &sync.Mutex{}

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	lock.Lock()
	myMap[n] = int64(res)
	lock.Unlock()
}

// 计算1-200的阶乘 存储在map中
func main() {

	for i := 1; i <= 200; i++ {
		go test(i)
	}

	time.Sleep(time.Second * 10)
	lock.Lock()
	for k, v := range myMap {
		fmt.Printf("myMap[%d]=%d\n", k, v)
	}
	lock.Unlock()
	// 主线程还没等goroutine执行完毕就退出了
}
