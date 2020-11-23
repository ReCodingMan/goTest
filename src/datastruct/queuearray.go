package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	maxSize int
	array [4]int //数组=>模拟队列
	front int //首位
	rear int //尾巴
}

//添加数据到队列
func (this *Queue) AddQueue(val int) (err error) {

	//先判断队列是否已满
	if this.rear == this.maxSize {
		return errors.New("queue full")
	}

	this.rear++
	this.array[this.rear] = val
	return
}

//取出队列
func (this *Queue) GetQueue() (val int, err error) {
	if this.front == this.rear {
		return -1, errors.New("queue empty")
	}

	this.front++
	val = this.array[this.front]
	return val, err
}

//显示队列，从队首遍历到队尾
func (this *Queue) ShowQueue() {
	fmt.Println("队列当前情况是：")
	for i:= this.front+1; i<=this.rear; i++ {
		fmt.Printf("array[%d]=%d\t", i, this.array[i])
	}
}

func main() {
	queue := &Queue{
		maxSize: 5,
		front: -1,
		rear: -1,
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
				err := queue.AddQueue(val)
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println("加入队列成功")
				}
			case "get":
				val, err := queue.GetQueue()
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println("取出了", val)
				}
			case "show":
				queue.ShowQueue()
			case "exit":
				break
			default:
				fmt.Println("输入有问题")
		}
	}
}
