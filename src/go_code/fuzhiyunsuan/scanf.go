package main 

import (
	"fmt"
)

func main() {
	var name string 
	var age int 
	var sal float32
	var isPass bool
	fmt.Println("请输入姓名， 年龄， 薪水， 是否通过 中间用逗号分隔开")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	fmt.Println("name = ", name, " age = ", age, " sal = ", sal, " isPass = ", isPass)
}