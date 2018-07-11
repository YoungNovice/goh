package main

import (
	"fmt"
)

/*
	go的运算符号

	1. 括号 ++ --
	2. 单目运算 + - ！ ~ type * & sizeof
	3. 算术 + - * / %
	4. 移位 << >>
	5. 关系 < <= > >= == !=
	6. 位 & ^ | 
	7。逻辑 && || !
	8. 赋值 = += -= *= /= %= >>= <<= &= ^= !=
	9. 逗号

*/

func main() {
	a := 3
	b := 4
	fmt.Println("a = ", a, " b = ", b)
	a = a + b // a = 7
	b = a - b // b = 3
	a = a - b // a = 4
	fmt.Println("a = ", a, " b = ", b)
	a, b = b, a
	fmt.Println("a = ", a, " b = ", b)
}