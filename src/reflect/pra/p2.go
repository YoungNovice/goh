package main 


import (
	"fmt"
	"reflect"
)

type Cal struct {
	Num1 int
	Num2 int
}

func (this Cal) GetSub(name string) int {
	return this.Num1 - this.Num2
}

func main () {
	c := Cal{
		Num1 : 8,
		Num2 : 3,
	}
	rType := reflect.TypeOf(c)
	rVal := reflect.ValueOf(c)

	for i := 0; i < rVal.NumField(); i ++ {
		rField := rType.Field(i)
		fmt.Println("字段的名称", rField.Name, " 字段的类型", rField.Type)

		fValField := rVal.Field(i)
		var kind reflect.Kind = fValField.Kind()
		fmt.Println("kind=", kind, "value=", fValField.Int())
	}

	method  := rVal.MethodByName("GetSub")
		var params []reflect.Value
		params = append(params, reflect.ValueOf("yangxuan"))
		result := method.Call(params)
		if result != nil {
			s := result[0].Int()
			fmt.Println("yangxuan", "完成了减法运行", s)
		}
}
