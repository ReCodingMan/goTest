package main

import "fmt"

func InsertSort(arr *[5]int) {
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

		fmt.Printf("第%d次插入后%v\n", j, *arr)
	}
}

func main() {
	//从大到小排序
	arr := [5]int{10, 34, 19, 100, 80}
	fmt.Println(arr)
	InsertSort(&arr)
}
