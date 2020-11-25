package main

import (
	"fmt"
	"math/rand"
	"time"
)

func InsertSort(arr *[80000]int) {
	for j:=1; j<len(arr); j++ {
		insertVal := arr[j]
		insertIndex := j-1	//下标

		//从大到小
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex + 1] = arr[insertIndex] //数据后移
			insertIndex--
		}

		//插入
		if insertIndex + 1 != j {
			arr[insertIndex + 1] = insertVal
		}

		//fmt.Printf("第%d次插入后%v\n", j, *arr)
	}
}

func main() {

	var arr [80000]int
	for i:=0; i<80000; i++ {
		rand.Seed(time.Now().UnixNano())
		arr[i] = rand.Intn(900000)
	}

	start := time.Now().Unix()
	InsertSort(&arr)
	end := time.Now().Unix()
	fmt.Printf("插入排序耗时=%d秒", end-start)

	//从大到小排序
	//arr := [5]int{10, 34, 19, 100, 80}
	//fmt.Println(arr)
	//InsertSort(&arr)
}
