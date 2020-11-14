package main

import (
	_ "fmt"
	"testing"
)

//编写一个测试用例，去测试 addUpper 是否正确
func TestGetSub(t *testing.T) {

	//调用
	res := getSub(10,3)
	if res != 7 {
		//fmt.Printf("AddUpper(7)执行错误，期望值=%v，实际值=%v\n", 55, res)
		t.Fatalf("getSub(7)执行错误，期望值=%v，实际值=%v\n", 7, res)
	}

	//如果正确，输出日志
	t.Logf("getSub(7)执行正确...")
}