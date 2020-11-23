package main

import (
	"errors"
	"fmt"
)

type CircleQueue struct {
	maxSize int
	array [5]int //数组=>模拟队列
	head int //首位
	tail int //尾巴
}

//添加数据到队列
func (this *CircleQueue) Push(val int) (err error) {
	if this.IsFull() {
		return errors.New("queue full")
	}

	this.array[this.tail] = val
	this.tail = (this.tail+1) % this.maxSize
	return
}

//取出队列
func (this *CircleQueue) Pop() (val int, err error) {
	if this.IsEmpty() {
		return 0, errors.New("queue empty")
	}

	val = this.array[this.head]
	this.head = (this.head+1) % this.maxSize
	return
}

//判断是否满了
func (this *CircleQueue) IsFull() bool {
	return (this.tail + 1) % this.maxSize == this.head
}

//判断是否为空
func (this *CircleQueue) IsEmpty() bool {
	return this.tail == this.head
}

//取出有多少个元素
func (this *CircleQueue) Size() int {
	return (this.tail + this.maxSize - this.head) % this.maxSize
}

func (this *CircleQueue) ListQueue() {
	size := this.Size()
	if size == 0 {
		fmt.Println("队列为空")
	}
	tempHead := this.head
	for i:=0; i<size; i++ {
		fmt.Printf("arr[%d]=%d\t", tempHead, this.array[tempHead])
		tempHead = (tempHead + 1) % this.maxSize
	}
	fmt.Println()
}

func main() {
	queue := &CircleQueue{
		maxSize: 5,
		head: 0,
		tail: 0,
	}
	var key string
	var val int
	for {
		fmt.Println("1.输入add")
		fmt.Println("2.输入get")
		fmt.Println("3.输入show")
		fmt.Println("4.输入exit")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("请输入要添加的数")
			fmt.Scanln(&val)
			err := queue.Push(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列成功")
			}
		case "get":
			val, err := queue.Pop()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("取出了", val)
			}
		case "show":
			queue.ListQueue()
		case "exit":
			break
		default:
			fmt.Println("输入有问题")
		}
	}
}