package main 

import (
	"fmt"
)

// 自己定义的空接口
type aa interface { }

func main () {

	var x aa
	var y = 10.0
	x = y
	switch i := x.(type) {
	case nil:
		fmt.Println("x 的类型： %T", i)
	case int:
		fmt.Println("x 是int")
	case float64:
		fmt.Println("x 是float64")
	case func(int) float64:
		fmt.Println("x 是func(int) float64型")
	case bool, string :
		fmt.Println("bool 或者 string")
	}
}