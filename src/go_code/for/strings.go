package main

import (
	"fmt"
)

func main() {

	// 传统的方式遍历 一个一个字节取数据
	str := "hello,world"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c \t", str[i])
	}

	str = "hello,world你好"
	strune := []rune(str)
	for i := 0; i < len(strune); i++ {
		fmt.Printf("%c \t", strune[i])
	}

	str1 := "ok,nihao哈哈"
	for index, value := range str1 {
		fmt.Printf("index=%d, value = %c\n", index, value)
	}		

}