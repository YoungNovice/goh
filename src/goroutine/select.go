package main 

import (
	"fmt"
	"strconv"
	"time"
)

// select处理从管道中取数据阻塞的问题
func main() {

	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan<-i
	}

	strChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		strChan<- "hello" + strconv.Itoa(i)
	}

	// 管道如果不关闭会导致阻塞 死锁
	label : for {
		select {
			case v := <-intChan:
			time.Sleep(time.Second)
				fmt.Println("intchan", v);
			case v := <-strChan:
			time.Sleep(time.Second)
				fmt.Println("strChan", v);
			default:
				fmt.Println("都取不到数据")
			time.Sleep(time.Second)
				break label
		}
	}
}
