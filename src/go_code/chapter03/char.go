package main 

// go 中字符用byte存储
// 字符串是一个个的字节组成
import (
	"fmt"
)

func main() {
	// byte = uint8
	var c1 byte = 'a'
	var c2 byte = '0'
	// 输出的是对应的ASCII
	fmt.Println("c1 = ", c1, "c2 = ", c2)
	// 如果需要输出字符需要格式化
	fmt.Printf("c1 = %c, c2 = %c\n", c1, c2) 

	// rune = int32
	var c3 rune = '北'
	fmt.Printf("c3 = %c 对应的值 = %d\n", c3, c3)

	// 字符本质是ASCII码值整型 运算的时候也是按照整型算的
}