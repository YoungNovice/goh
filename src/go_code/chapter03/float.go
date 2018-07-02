package main 

import (
	"fmt"
)

func main() {
	var price  = 89.1
	var num1 float32 = 1.1
	var num2 float64 = 1.2

	fmt.Printf("price type %T, num1 type %T num2 type %T\n", price, num1, num2)

	num6 := 1.3
	// 0.123
	num7 := .123
	fmt.Println("num6=", num6, "num7=", num7)
	
	// 浮点数的科学计数法 5.12 * 10 ^ 2
	num8 := 5.12e2 
	num9 := 5.12E2
	// 5.12 * 10 ^ -2
	num10 := 5.12E-2 

	fmt.Println("num8=", num8, "num9=", num9, "num10=", num10)
}