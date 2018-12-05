package main 

import (
	"fmt"
)


func writeData(intchan chan int) {
	for i := 1; i<= 50; i++ {
		intchan <- i
		fmt.Println("writedata i= ", i)
	}
	// 不close会造成死锁
	close(intchan)
}

func readData(intchan chan int, exitchan chan bool) {
	for {
		v, ok := <-intchan
		if !ok {
			fmt.Println("本次没有读取到数据")
			break
		}
		fmt.Println("readdata v= ", v)
	}
	// 向退出管道写数据	
	exitchan <- true
	close(exitchan)
}

// 这种chan需要自己搞个死循环来阻塞 太sb了 
// 而且如果不close的读会死锁的
func main() {
	intchan := make(chan int, 50)
	extichan := make(chan bool, 1)
	i := 1
	
	go writeData(intchan)
	go readData(intchan, extichan)
	for {
		i++
		_, ok := <-extichan
		if !ok {
			fmt.Println("boolchan 中没有数据")
			break
		}
		fmt.Println("exitchan 主线程中for循环执行了 ", i)
	}

}
