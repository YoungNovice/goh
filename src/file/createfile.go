package main 

import (
	"fmt"
	"os"
	"bufio"
)
// 创建一个新文件 写入hello go
// 打开一个存在的文件 将原来的内容覆盖成 你好xxx
// 打开一个存在的文件 在原来的内容追加内容 abc
// 打开一个存在的文件 将原来的内容显示在终端 并且追加 5句北京你好

func main() {
	file, err := os.OpenFile("/Users/yangxuan/openfile.md", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("打开文件失败")
		return
	} else {
		defer file.Close()
	}
	writer := bufio.NewWriter(file)
	for i := 0; i<5; i++ {
		writer.WriteString("hello go\n")
	}
	writer.Flush()
}
