package main 

import (
	"fmt"
	"strconv"
)

// 基本类型根string相互转
// 重点 
func main() {
	// typeToString()
	typeToString2()
}

// 基本类型转string 可用fmt包下面的Sprintf函数实现
func typeToString() {
	var num1 int = 99
	var num2 float64 = 23.1
	var b bool = true
	var myChar byte = 'h'


	str1 := fmt.Sprintf("%d", num1)
	str2 := fmt.Sprintf("%f", num2)
	str3 := fmt.Sprintf("%t", b)
	str4 := fmt.Sprintf("%b", myChar)
	fmt.Printf("type %T = %q\n", str1, str1)
	fmt.Printf("type %T = %q\n", str2, str2)
	fmt.Printf("type %T = %q\n", str3, str3)
	fmt.Printf("type %T = %q\n", str4, str4)
}

// 用strconv包下面的函数实现
/*
	func Formatbool(b bool) string
	func FormatFloat(f float64, byte, prec , bitSize int) strin
	func FormatInt(i int64, base int) string
	func FormatUint(i int64, base int) string
*/
func typeToString2() {
	var num1 int = 33
	var num2 float64 = 24.1
	var b bool = false
	var myChar byte = 'h'
	// 10进制
	str1 := strconv.FormatInt(int64(num1), 10)
	// f格式  小数保留10位 这个数字是float64
	str2 := strconv.FormatFloat(num2, 'f', 10, 64)
	str3 := strconv.FormatBool(b)
	str4 := strconv.FormatUint(uint64(myChar), 10)
	fmt.Printf("type %T = %q\n", str1, str1)
	fmt.Printf("type %T = %q\n", str2, str2)
	fmt.Printf("type %T = %q\n", str3, str3)
	fmt.Printf("type %T = %q\n", str4, str4)
}