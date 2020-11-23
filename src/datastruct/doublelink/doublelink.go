package main

import "fmt"

//定义一个heroNode
type HeroNode struct {
	no int
	name string
	nickname string
	pre *HeroNode
	next *HeroNode
}

//给双向链表插入节点
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
	newHeroNode.pre = temp
}

//按编号从小到大排序
func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	//先找到该链表最后节点
	temp := head
	flag := true
	for {
		if temp.next == nil {
			break
		} else if temp.next.no > newHeroNode.no {
			//找到最后
			break
		} else if temp.next.no == newHeroNode.no {
			flag = false
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Println("已经存在，无法插入...")
		return
	}
	newHeroNode.next = temp.next
	temp.next = newHeroNode
}

func DelHeroNode(head *HeroNode, id int) {
	temp := head
	flag := false
	for {
		if temp.next == nil {
			break
		} else if temp.next.no == id {
			flag = true
			//找到最后
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Println("没找到，没法删除...")
		return
	}
	//删除
	temp.next = temp.next.next
}

//显示链表的所有信息
//单向双向一样
func ListHeroNode(head *HeroNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("空链表...")
	}
	//倒叙打印
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}

	for {
		fmt.Printf("[%d, %s, %s]==>", temp.no, temp.name, temp.nickname)
		temp = temp.pre
		if temp.pre == nil {
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
		no: 4,
		name: "卢俊义",
		nickname: "玉麒麟",
	}
	hero3 := &HeroNode{
		no: 2,
		name: "林冲",
		nickname: "豹子头",
	}

	//3、加入
	InsertHeroNode(head, hero1)
	InsertHeroNode(head, hero2)
	InsertHeroNode(head, hero3)
	//DelHeroNode(head,4)
	//DelHeroNode(head,1)
	//DelHeroNode(head,3)
	//4、显示
	ListHeroNode(head)
}