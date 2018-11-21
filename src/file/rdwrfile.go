package main 

import (
	"fmt"
	"os"
	"bufio"
	"io"
)
// 打开一个存在的文件 将原来的内容显示在终端 并且追加 5句北京你好

func main() {
	file, err := os.OpenFile("/Users/yangxuan/openfile.md",os.O_RDWR| os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("打开文件失败")
		return
	} else {
		defer file.Close()
	}
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { 
			break
		}
		fmt.Print(str)
	}
	writer := bufio.NewWriter(file)
	for i := 0; i<5; i++ {
		writer.WriteString("读写的方式打开\r\n")
	}
	writer.Flush()
}
