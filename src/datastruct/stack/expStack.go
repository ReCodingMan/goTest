package main

import (
	"fmt"
)

//判断一个字符是否是一个运算符
func (this *Stack) IsOper(val int) bool {
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	} else {
		return false
	}
}

//运算方法
func (this *Stack) Cal(num1 int, num2 int, oper int) int {
	res := 0
	switch oper {
		case 42:
			res = num2 * num1
		case 43:
			res = num2 + num1
		case 45:
			res = num2 - num1
		case 47:
			res = num2 / num1
		default:
			fmt.Println("运算符有问题")
	}
	return res
}

//编写一个方法，返回优先级。【*/ = 1】，【+- = 0】
func (this *Stack) Priority(oper int) int {
	res := -1
	if oper == 42 || oper == 47 {
		res = 1
	} else if oper == 43 || oper == 45 {
		res = 0
	}
	return res
}

func main() {

	//数栈
	numStack := &Stack{
		MaxTop: 20,
		Top: -1,
	}

	//符号栈
	operStack := &Stack{
		MaxTop: 20,
		Top: -1,
	}

	exp := "3+2*6-2"
	//定义 index，帮助扫描 exp
	for {
		ch := []
	}
}