package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Stack struct {
	MaxTop int	//栈最大存放个数
	Top int	//栈顶
	arr [20]int
}

func (this *Stack) Push(val int) (err error) {

	//判断栈是否满了
	if this.Top == this.MaxTop-1 {
		fmt.Println("栈满了")
		return errors.New("stack full")
	}

	this.Top++
	//放入数据
	this.arr[this.Top] = val
	return
}

func (this *Stack) Pop() (val int, err error) {
	if this.Top == -1 {
		fmt.Println("栈空")
		return 0, errors.New("栈空")
	}
	val = this.arr[this.Top]
	this.Top--
	return val, nil
}

//遍历栈，从栈底开始
func (this *Stack) List() {
	if this.Top == -1 {
		fmt.Println("栈空")
		return
	}

	for i:=this.Top; i>=0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, this.arr[i])
	}
}

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
	index := 0
	num1 := 0
	num2 := 0
	oper := 0
	result := 0

	for {
		ch := exp[index:index+1] //字符串
		temp := int([]byte(ch)[0]) //转成切片，取第一个。就是字符对应的 ascii
		//fmt.Println(temp)
		if operStack.IsOper(temp) {

			//符号，分两个逻辑
			if operStack.Top == -1 {
				//空栈
				operStack.Push(temp)
			} else {
				//判断优先级
				if operStack.Priority(operStack.arr[operStack.Top]) >= operStack.Priority(temp) {
					num1,_ = numStack.Pop()
					num2,_ = numStack.Pop()
					oper,_ = operStack.Pop()
					result = operStack.Cal(num1, num2, oper)

					numStack.Push(result)
					operStack.Push(temp)
				} else {
					operStack.Push(temp)
				}
			}

		} else {
			val ,_ := strconv.ParseInt(ch, 10, 64)
			numStack.Push(int(val))
		}

		if index+1 == len(exp) {
			break
		}
		index++
	}

	//开始计算
	for {
		if operStack.Top == -1 {
			break
		}

		num1,_ = numStack.Pop()
		num2,_ = numStack.Pop()
		oper,_ = operStack.Pop()
		result = operStack.Cal(num1, num2, oper)

		numStack.Push(result)

	}

	//如果算法没有问题，结果就是最后的数
	res,_ := numStack.Pop()
	fmt.Printf("表达式为%s = %v", exp, res)
}