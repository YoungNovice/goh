package main

import (
	"fmt"
	"reflect"
)

func testInt(a interface{}) {
	// 获取Type
	rType := reflect.TypeOf(a)
	fmt.Println("type = ", rType)
	
	rVal := reflect.ValueOf(a)
	iv := rVal.Interface()
	fmt.Println("value =", rVal)
	fmt.Println("Interface =", iv)
	v := rVal.Int() 
	fmt.Println("v =", v)
	res, ok := iv.(int)
	if ok {
		fmt.Println("真实的值是=", res)
	}
}

func main() {
	var a int = 10
	testInt(a)


}
