package main 

import (
	"fmt"
	"time"
	"strconv"
)


func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("test() hello, world", strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {

	// 开启一个goroutine
	go test()

	for i := 1; i <= 10; i++ {
		fmt.Println("main() hello, main", strconv.Itoa(i))
		time.Sleep(time.Second)
	}

}
