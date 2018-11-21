package main 

import (
	"io/ioutil"
	"fmt"
)

func main() {
	content, err := ioutil.ReadFile("/Users/yangxuan/a.txt")
	if err == nil {
		er := ioutil.WriteFile("/Users/yangxuan/b.txt", content, 0666)
		if er != nil {
			fmt.Println("写入阶段失败")
		}
	} else {
		fmt.Println("读入阶段失败")
	}


}
