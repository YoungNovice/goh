package main

import (
	"fmt"
	"reflect"
)

func main() {
	var v float64 = 1.2

	rVal := reflect.ValueOf(v)
	rType := rVal.Type()
	fmt.Printf("rtype=%v, kind=%v\n", rType, rVal.Kind())

	iv := rVal.Interface()

	v2, ok := iv.(float64)
	if ok {
		fmt.Println(v2)
	}

}
