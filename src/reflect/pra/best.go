package main 

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Score int 
	Sex string
}

func (this Monster) Print() {
	fmt.Println("method print", this.Name, this.Age)
}

func (this Monster) GetSum(a, b int) int {
	return a + b
}

func (this *Monster) Set(name string, age int) {
	this.Name = name
	this.Age = age
}

func ref(a interface{}) {
	 rType := reflect.TypeOf(a)
	 rVal := reflect.ValueOf(a)

	if rType.Kind() != reflect.Struct {
		fmt.Println("不是结构体")
		return 
	}

	numF := rVal.NumField()

	for i := 0; i < numF; i ++ {
		var tField reflect.StructField = rType.Field(i)
		fmt.Printf("字段名称%s, 字段的类型是%v, 字段的jsonTag=%v\n", tField.Name, tField.Type, tField.Tag.Get("json"))
		rField := rVal.Field(i)
		fmt.Println("rField= ", rField)
	}

	numM := rVal.NumMethod()
	fmt.Printf("一共有%d个方法\n", numM)

	// 调用第二个方法 按照函数名字排序的
	rVal.Method(1).Call(nil)
}

func main() {
	m := Monster{
		Name : "string", 
		Age : 18,
	}
	ref(m)
	m.Set("aa", 2)
	fmt.Println(m)



}
