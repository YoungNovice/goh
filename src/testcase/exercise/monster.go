package monster

import (
	"encoding/json"
	"os"
	"bufio"
	"io/ioutil"
	"fmt"
)

type Monster struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Skill string `json:"skill"`
}

// 将Monster序列化并保存在文件中	
func (m *Monster) Store() error {
	bytes, marshalErr := json.Marshal(m)
	if marshalErr != nil {
		fmt.Println("序列化Monster结构体失败")
		return marshalErr
	}
	result := string(bytes)
	fileName := "/Users/yangxuan/monster.json"
	file, fileErr := os.OpenFile(fileName, os.O_RDWR | os.O_CREATE | os.O_TRUNC, 0777)
	if fileErr != nil {
		fmt.Println("创建文件失败")
		return fileErr
	} else {
		defer file.Close();
	}
	writer := bufio.NewWriter(file)
	_, writeErr := writer.WriteString(result)
	if writeErr != nil {
		fmt.Println("写入json字符串失败")
		return writeErr
	}
	writer.Flush()
	return nil
}

// filename json文件的地址
func (m *Monster) Restore(filename string) error {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		// 读取失败
		return err
	}
	jsonerr := json.Unmarshal(contents, m)
	if jsonerr != nil {
		// 反序列化失败
		return jsonerr
	}
	return nil
}
