package main

import (
	"fmt"
	"os"
)

//定义emp
type Emp struct {
	Id int
	Name string
	Next *Emp
}

//显示自己
func (this *Emp) ShowMe() {
	fmt.Printf("链表%d 找到雇员%d \n", this.Id % 7, this.Id)
}

//定义EmpLink
//这里不带表头，第一个就存放雇员
type EmpLink struct {
	Head *Emp
}

//添加员工方法，编号从小到大
func (this *EmpLink) Insert(emp *Emp) {

	cur := this.Head //这是辅助指针
	var pre *Emp = nil //这是一个辅助指针，在cur前面
	if cur == nil {
		this.Head = emp
		return
	} else {
		//如果不是空链表
		for {
			if cur != nil {
				//比较
				if cur.Id > emp.Id {
					break
				}
				pre = cur
				cur = cur.Next
			} else {
				break
			}
		}

		pre.Next = emp
		emp.Next = cur
	}
}

//显示当前链表信息
func (this *EmpLink) ShowLink(no int) {
	if this.Head == nil {
		fmt.Printf("链表%d为空\n", no)
		return
	}

	//遍历当前链表，并显示数据
	cur := this.Head
	for {
		if cur != nil {
			fmt.Printf("链表%d 雇员id=%d 名字=%s", no, cur.Id, cur.Name)
			cur = cur.Next
		} else {
			break
		}
	}
	fmt.Println()
}

//根据id查找对应雇员，如果没有就返回nil
func (this *EmpLink) FindById(id int) *Emp {
	cur := this.Head
	for {
		if cur != nil && cur.Id == id {
			return cur
		} else if cur == nil {
			break
		}
		cur = cur.Next
	}
	return nil
}

//定义hashtable，含有一个链表数组
type HashTable struct {
	LinkArr [7]EmpLink
}

//给 HashTable 添加 Insert 雇员的方法
func (this *HashTable) Insert(emp *Emp) {
	//使用散列函数，确定将雇员添加到哪个链表
	linkNo := this.HashFun(emp.Id)
	//使用对应的链表添加
	this.LinkArr[linkNo].Insert(emp)
}

//编写散列方法
func (this *HashTable) HashFun(id int) int {
	return id % 7 //得到链表的下标
}

//编写一个方法，查找
func (this *HashTable) FindById(id int) *Emp {
	//找出在哪个链表
	linkNo := this.HashFun(id)
	return this.LinkArr[linkNo].FindById(id)
}

//显示hashtable所有雇员
func (this *HashTable) ShowAll() {
	for i:=0; i<len(this.LinkArr); i++ {
		this.LinkArr[i].ShowLink(i)
	}
}

func main() {
	key := ""
	id := 0
	name := ""
	var hashtable HashTable

	for {
		fmt.Println("================雇员系统菜单==============")
		fmt.Println("input 表示添加雇员")
		fmt.Println("show 表示显示雇员")
		fmt.Println("find 表示查找雇员")
		fmt.Println("exit 表示退出系统")
		fmt.Println("请输入你的选择")
		fmt.Scanln(&key)
		switch key {
		case "input":
			fmt.Println("输入雇员id")
			fmt.Scanln(&id)
			fmt.Println("输入雇员name")
			fmt.Scanln(&name)
			emp := &Emp{
				Id: id,
				Name: name,
			}
			hashtable.Insert(emp)
		case "show":
			hashtable.ShowAll()
		case "find":
			fmt.Println("请输入id号：")
			fmt.Scanln(&id)
			emp := hashtable.FindById(id)
			if emp == nil {
				fmt.Printf("id=%d 的雇员不存在", id)
			} else {
				//编写一个方法，显示雇员信息
				emp.ShowMe()
			}
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("输入有误")
		}
	}
}
