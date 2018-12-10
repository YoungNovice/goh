package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	age int
	Sex int
}

func main() {
	stu := Student{
		Name : "Yangxuan",
		age :15,
		Sex : 1,
	}
	var a interface{} = stu
	rType := reflect.TypeOf(a)
	fmt.Println("rType=", rType)

	rVal := reflect.ValueOf(a)
	fmt.Println("rVal=", rVal)

	itf := rVal.Interface()

	switch  v :=itf.(type) {
	case Student:
		fmt.Println("stu.Name=%s, stu.age=%d, stu.Sex=%d", v.Name, v.age, v.Sex)
	}
}
