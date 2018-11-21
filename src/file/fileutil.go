package main 

// 一次性读入文件
import (
	"fmt"
	"io/ioutil"
)
func main() {
	content, err := ioutil.ReadFile("/Users/yangxuan/unzip.org")
	if err != nil {
		fmt.Printf("read file error=%v", err)
	}

	fmt.Printf("%v", string(content))

}
