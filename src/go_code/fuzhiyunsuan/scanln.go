package main 

import (
	"fmt"
)

// 函数scanln 逐行读取标准输入流中的数据
func main () {
	var name string 
	var age byte
	var sal float32
	var isPass bool
	fmt.Println("请输入姓名")
	fmt.Scanln(&name)
	fmt.Println("请输入年龄")
	fmt.Scanln(&age)
	fmt.Println("请输入薪水")
	fmt.Scanln(&sal)
	fmt.Println("请输入是否通过")
	fmt.Scanln(&isPass)

	fmt.Println("name = ", name, " age = ", age, " sal = ", sal, " isPass = ", isPass)

	 
}