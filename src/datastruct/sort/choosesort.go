package main

import "fmt"

func SelectSort(arr *[5]int) {
	for j:=0; j<len(arr)-1; j++ {
		max := arr[j]
		maxIndex := j

		//遍历后面的数
		for i:=j+1; i<len(arr); i++ {
			if max < arr[i] {
				max = arr[i]
				maxIndex = i
			}
		}
		//交换
		if maxIndex != j {
			arr[j], arr[maxIndex] = arr[maxIndex], arr[j]
		}

		fmt.Printf("第%d次结果%v\n", j, arr)
	}
}

func main() {
	//从大到小排序
	arr := [5]int{10, 34, 19, 100, 80}
	fmt.Println(arr)
	SelectSort(&arr)
}
