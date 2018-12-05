package main 

import (
	 "fmt"
)

type Cat struct {
	Name string
}
func main() {

	var allchan chan interface{}
	allchan = make(chan interface{}, 4)
	cat1 := Cat{Name :"y"}
	cat2 := Cat{Name :"z"}
	allchan <- cat1
	allchan <- cat2
	allchan <- 1
	allchan <- "jack"

	xxx := <- allchan
	res, ok := xxx.(Cat)
	if ok {
		fmt.Println(res.Name)	
	}
}
