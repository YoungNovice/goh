package main 
// 基本类型的相互转换
// 所有的类型都需要相互转换

import (
	"fmt"
)

func main () {
	var i int32 = 100
	var n1 float32 = float32(i)
	fmt.Println("i=%v, n1=%v", i, n1)
}
