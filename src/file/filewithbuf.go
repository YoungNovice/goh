package main 

// 数据在数据源和内存之间经历的路径
// 输入流 从数据源到内存
// 输出流 从内存到数据源
import (
	"fmt"
	"os"
	"bufio"
	"io"
)

// file 是*File


func main() {
	file, err := os.Open("/Users/yangxuan/unzip.org")
	if err != nil {
		fmt.Println("打开文件失败")
	} else {
		defer file.Close()
	}
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break;
		}
		fmt.Print(str)
	}


}
