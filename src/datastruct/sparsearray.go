package main

import (
	"fmt"
)

type ValNode struct {
	row int
	col int
	val int
}

func main() {
	//1、先创建原始数组
	var chessMap[11][11]int
	chessMap[1][2] = 1//黑子
	chessMap[2][3] = 2//蓝子

	//2、输出原始数组
	for _,v := range chessMap {
		for _,v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

	//3、转成稀疏数组
	var sparseArr []ValNode
	//初始化一个节点
	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr, valNode)

	for i,v := range chessMap {
		for j,v2 := range v {
			if v2 != 0 {
				//创建一个节点
				valNode = ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
		fmt.Println()
	}

	//4、输出稀疏数组
	for i, valNode := range sparseArr {
		fmt.Printf("%d: %d %d %d\n", i, valNode.row, valNode.col, valNode.val)
	}

	//5、恢复原始数组
	var chessMap2 [11][11]int
	//遍历稀疏数组/文件每一行
	for i, valNode := range sparseArr {
		if i == 0 {
			continue
		}
		chessMap2[valNode.row][valNode.col] = valNode.val
	}
	for _,v := range chessMap2 {
		for _,v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

}
