package main

import "fmt"

// 小写的整个包中都有效果
var age int = 50

// 大写的整个程序中都有效果
var Name string = "jack"

func test() {
	// age 和 name 的变量的作用域只在test函数内部
	age := 10
	Name := "tom~"
	fmt.Println("age = ", age)
	fmt.Println("Name = ", Name)

}

func main() {
	fmt.Println("main age = ", age)
	fmt.Println("main Name = ", Name)
	test()
}
