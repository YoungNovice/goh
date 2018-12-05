package main

// var name chan datatype
// var intChan chan int
// var mapChan chan map[int]string
// var perChan chan Person
// var perChan2 chan *Person

// chan 是引用类型的 必须make后才能使用 intChan只能写入int类型的数据
// chan本质上是一个线程操作安全的队列

import (
	"fmt"
)

func main () {
	// 创建一个3个容量的int队列
	var intChan chan int = make(chan int, 3)

	fmt.Printf("地址是%v, &intchan=%p\n", intChan, &intChan)

	// 3向管道中写入数据10
	intChan <- 10
	num := 200
	intChan <- num

	// 管道的len 和 cap容量
	fmt.Printf("intChan len=%v, intChan cap=%v\n", len(intChan), cap(intChan))

	// 从管道中取数据
	var num2 int
	num2 = <- intChan
	fmt.Println("num2=", num2)
}
