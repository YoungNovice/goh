package main

import (
	"fmt"
)

// init 被自动调用
func init () {
	fmt.Println("init")
}
// 没个源文件都可以包含一个init函数， 该函数
// 会在main函数执行前 被go运行框架， 也就是
// 说init函数会在main函数前被调用
func main() {
	fmt.Println("main")
}
