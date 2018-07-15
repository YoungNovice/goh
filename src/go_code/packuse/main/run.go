package main 
// 结论一 同一个文件夹下面只能又一个包名
// 同一个包下不能有重复东西
import (
	"go_code/packuse/util"
)

// 编译项目:
// cd 到project跟目录
// 编译main包输出到bin下
// go build -o bin/aaa.exe go_code/packuse/main

func main() {
	utils.Abb()
}