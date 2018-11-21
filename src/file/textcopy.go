package main

import (
	f "fmt"
	"io"
	"os"
)

func main() {
	w, err := CopyFile("/Users/yangxuan/f.mp4", "/Users/yangxuan/3.mp4")
	if err != nil {
		f.Println(err.Error())
	}
	f.Println(w)
}

func CopyFile(src, des string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		f.Println(err)
	}
	defer srcFile.Close()

	desFile, err := os.Create(des)
	if err != nil {
		f.Println(err)
	}
	defer desFile.Close()

	return io.Copy(desFile, srcFile)
}
