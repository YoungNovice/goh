package main 

import (
	"fmt"
	"strings"
)

// 用一个闭包给文件名添加后缀 如果有就不用添加了
func makeSuffix(suffix string) func(string) string {
	return func (filename string) string {
		if (strings.HasSuffix(filename, suffix)) {
			return filename
		}
		return filename + suffix
	}
}

func main() {
	f := makeSuffix(".jpg")
	fmt.Println(f("aa.jpgdf"))
}