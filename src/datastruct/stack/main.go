package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	MaxTop int	//栈最大存放个数
	Top int	//栈顶
	arr [5]int
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

func main() {

	stack := Stack{
		MaxTop: 5,
		Top: -1, //表示栈为空
	}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.List()

	val, _ := stack.Pop()
	fmt.Println("出栈", val)
	stack.List()

}
