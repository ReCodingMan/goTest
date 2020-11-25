package main

import "fmt"

//编写一个函数，找出路
func SetWay(myMap *[8][7]int, i int, j int) bool {

	//找到出口
	if myMap[6][5] == 2 {
		return true
	} else {
		//继续找
		if myMap[i][j] == 0 {
			//可以探测，假设这个点通的，但是需要探测。优先级：上下左右
			myMap[i][j] = 2
			if SetWay(myMap, i-1, j) {//上
				return true
			} else if SetWay(myMap, i+1, j) {//下
				return true
			} else if SetWay(myMap, i, j-1) {//左
				return true
			} else if SetWay(myMap, i, j+1) {//右
				return true
			} else {//死路一条
				myMap[i][j] = 3
				return false
			}
		} else {
			//不可以探测
			return false
		}
	}
}

func main() {

	//1、创建一个二维数组，模拟迷宫
	/**
		规则：
			1、如果元素为1，是墙
			2、如果元素为0，没有走过的点
			3、如果元素为2，是一个通路
			4、如果元素为3，走过的，但是走不通
	 */
	var myMap [8][7]int
	for i:=0; i<7; i++ {
		myMap[0][i] = 1
		myMap[7][i] = 1
	}
	for i:=0; i<8; i++ {
		myMap[i][0] = 1
		myMap[i][6] = 1
	}
	myMap[3][1] = 1
	myMap[3][2] = 1

	//打印
	for i:=0; i<8; i++ {
		for j:=0; j<7; j++ {
			fmt.Print(myMap[i][j], " ")
		}
		fmt.Println()
	}

	//使用测试
	SetWay(&myMap, 1, 1)

	//打印
	for i:=0; i<8; i++ {
		for j:=0; j<7; j++ {
			fmt.Print(myMap[i][j], " ")
		}
		fmt.Println()
	}
}
