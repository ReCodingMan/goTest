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

//删除一只猫
func DelCatNode(head *CatNode, id int) *CatNode {
	temp := head
	helper := head
	if temp.next == nil {
		fmt.Println("空的环形链表")
		return head
	}

	//如果只有一个节点
	if temp.next == head {
		temp.next = nil
		return head
	}

	//helper 指到最后
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}

	//有多个节点
	flag := true
	for {
		if temp.next == head {
			//找了一圈了
			break
		}
		if temp.no == id {
			//找到了
			if temp == head {
				head = head.next
			}
			helper.next = temp.next
			flag = false
			break
		}
		temp = temp.next
		helper = helper.next
	}

	if flag {
		//这里比较最后一次
		if temp.no == id {
			helper.next = temp.next
		}
	}
	return head
}

func main() {

	//初始化一个头节点
	head := &CatNode{}

	//创建一只猫
	cat1 := &CatNode{
		no: 1,
		name: "tom",
	}
	cat2 := &CatNode{
		no: 2,
		name: "tom",
	}
	cat3 := &CatNode{
		no: 3,
		name: "tom",
	}

	InsertCatNode(head, cat1)
	InsertCatNode(head, cat2)
	InsertCatNode(head, cat3)
	head = DelCatNode(head, 1)
	ListCircleLink(head)
}