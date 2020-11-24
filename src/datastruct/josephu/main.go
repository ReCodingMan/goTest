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
	fmt.Println()
}

//首指针，从第几个开水，报数
func PlayGame(first *Boy, startNo int, countNum int) {

	if first.Next == nil {
		fmt.Println("空的链表")
		return
	}

	//开始 < 小孩数量

	tail := first
	//将 tail 指到最后，方便后面删除
	for {
		if tail.Next == first {
			break
		}
		tail = tail.Next
	}

	//让 first 移动到 startNo
	for i:=1; i<=startNo-1; i++ {
		first = first.Next
		tail = tail.Next
	}

	//开始数数，开始删除
	for {
		for i:=1; i<=countNum-1; i++ {
			first = first.Next
			tail = tail.Next
		}
		fmt.Printf("编号为%d出列 ->\n", first.No)
		//删除 first 节点
		first = first.Next
		tail.Next = first

		//退出条件
		if tail == first {
			break
		}
	}

	fmt.Printf("最后的编号为%d", first.No)
}

func main() {
	first := AddBoy(5)
	ShowBoy(first)
	PlayGame(first, 2, 3)
}
