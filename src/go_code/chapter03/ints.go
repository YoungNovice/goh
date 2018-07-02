package main 

import (
	"fmt"
	"unsafe"
)

func main() {

	var n = 100
	fmt.Printf("n 的类型是： %T\n", n)

	var n2 int64 = 10
	// unsafe包中的Sizeof可以查看数据类型的占用的字节数
	fmt.Printf("n2 的类型是 %T n2占用的字节数是 %d\n", n2, unsafe.Sizeof(n2))
}