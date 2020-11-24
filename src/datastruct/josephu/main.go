package main

import "fmt"

type Boy struct {
	No int
	Next *Boy
}

//编写一个函数，构成单向的环形列表，n个个数，返回头指针
func AddBoy(n int) *Boy {
	first := &Boy{}
	curBoy := &Boy{}

	if n < 1 {
		fmt.Println("n太小了")
		return first
	}

	//循环构建循环列表
	for i:=1; i<=n; i++ {
		boy := &Boy{
			No: i,
		}
		//需要一个辅助指针

		if i==1 {
			first = boy
			curBoy = boy
			curBoy.Next = first //
		} else {
			curBoy.Next = boy
			//boy.Next = first
			curBoy = boy
			curBoy.Next = first
		}
	}

	return first
}

func ShowBoy(first *Boy) {

	if first.Next == nil {
		fmt.Println("空链表")
		return
	}

	//创建一个指针，帮助遍历
	curBoy := first
	for {
		fmt.Printf("小孩编号=%d ->", curBoy.No)
		if curBoy.Next == first {
			break
		}
		curBoy = curBoy.Next
	}
}

func main() {
	first := AddBoy(5)
	ShowBoy(first)
}
