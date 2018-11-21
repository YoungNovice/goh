package main 

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(dest string, src string) (written int64, err error) {
	file, err := os.Open(src)
	if err != nil {
		fmt.Println("获取源文件失败")
		return 0, err
	}
	defer file.Close()
	// 读取Reader
//	reader := bufio.NewReader(file)

	filew, errw := os.Create(dest)
	if errw != nil {
		fmt.Println("获取目标文件失败")
		return 0, errw
	}
	defer filew.Close()
	// writer
//	writer := bufio.NewWriter(filew)

	return io.Copy(filew, file)
}
func main() {
	src := "/Users/yangxuan/1.mp4";
	dest := "/Users/yangxuan/b.mp4";
	written, err := CopyFile(dest, src)
	if err != nil {
		fmt.Println("复制文件失败")
		return 
	}
	fmt.Printf("一共复制了%v个字节", written)
}
