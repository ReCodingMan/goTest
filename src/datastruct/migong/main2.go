package main

import "fmt"

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
	
	//打印
	for i:=0; i<8; i++ {
		for j:=0; j<7; j++ {
			fmt.Print(myMap[i][j], " ")
		}
		fmt.Println()
	}

}
