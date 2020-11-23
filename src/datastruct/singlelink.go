package main

import "fmt"

//定义一个heroNode
type HeroNode struct {
	no int
	name string
	nickname string
	next *HeroNode
}

//给链表插入节点
//在最后插入
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	//先找到该链表最后节点
	temp := head
	for {
		if temp.next == nil {
			//找到最后
			break
		}
		temp = temp.next//不断指向下一个节点
	}

	temp.next = newHeroNode
}

//显示链表的所有信息
func ListHeroNode(head *HeroNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("空链表...")
	}

	for {
		fmt.Printf("[%d, %s, %s]==>", temp.next.no, temp.next.name, temp.next.nickname)
		temp = temp.next
		if temp.next == nil {
			fmt.Println("打印结束")
			break
		}
	}
}

func main() {

	//1、创建一个头节点
	head := &HeroNode{}

	//2、创建一个新的 HeroNode
	hero1 := &HeroNode{
		no: 1,
		name: "宋江",
		nickname: "及时雨",
	}
	hero2 := &HeroNode{
		no: 2,
		name: "卢俊义",
		nickname: "玉麒麟",
	}

	//3、加入
	InsertHeroNode(head, hero1)
	InsertHeroNode(head, hero2)
	//4、显示
	ListHeroNode(head)
}