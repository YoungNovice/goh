package main 


import (
	"fmt"
	"reflect"
)

func main() {
	var a int = 10
	// 修改值还必须得用地址去弄
	// 但是地址的Kind是ptr 不是int 所以不能直接调用SetInt 得获取其element
	rVal := reflect.ValueOf(&a)
	rVal.Elem().SetInt(20)

	fmt.Println(a)
}


