package util

import (
	"testing"
)
// command line
// go test -v
// go test -v cal_test.go cal.go 测试单个文件
// go test -v -test.run funcname 测试单个函数

// 测试用例必须Test开头 并且Test后面的第一个字符必须是大写
func TestAddUpper(t *testing.T) {
	res := addUpper(10)
	if res != 55 {
		// 在调用Logf之后调用FailNow
		t.Fatalf("addUpper(10)执行错误， 期望值%v, 实际值", 55, res)
	}
	t.Logf("addUpper执行正确")
}

func TestGetSub(t *testing.T) {
	res := getSub(10, 8)
	if res != 2 {
		t.Fatalf("getsub(10, 8)执行错误， 期望值%v, 实际值", 2, res)
	}
	t.Logf("getSub执行正确")
}
