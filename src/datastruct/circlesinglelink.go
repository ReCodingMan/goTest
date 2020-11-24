package main

import "fmt"

type CatNode struct {
	no int
	name string
	next *CatNode
}

func InsertCatNode(head *CatNode, newCatNode *CatNode) {
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head //形成一个环状
		return
	}

	//定义一个临时变量，找到最后一个节点
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	temp.next = newCatNode
	newCatNode.next = head
}

func ListCircleLink(head *CatNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("空空如也~")
		return
	}

	for {
		fmt.Printf("猫的信息=[id=%d name=%s] -> \n", temp.no, temp.name)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

func main() {

	//初始化一个头节点
	head := &CatNode{}

	//创建一只猫
	cat1 := &CatNode{
		no: 1,
		name: "tom",
	}

	InsertCatNode(head, cat1)
	ListCircleLink(head)
}