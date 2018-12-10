package main

import (
	"fmt"
	"reflect"
)


type myInt int

func main() {

	var a myInt = 1
	fmt.Println(a)
	rType := reflect.TypeOf(a)
	kind := rType.Kind()
	fmt.Println(rType, kind)
}
