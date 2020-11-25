package main

import (
	"fmt"
	"math/rand"
	"time"
)

func QuickSort(left int, right int, array *[80000]int) {
	l := left
	r := right

	pivot := array[(left + right)/2]
	for ; l<r; {
		for ; array[l] < pivot; {
			l++
		}
		for ; array[r] > pivot; {
			r--
		}
		if l >= r {
			break
		}

		array[l], array[r] = array[r], array[l]

		if array[l] == pivot {
			l++
		}
		if array[r] == pivot {
			r--
		}
	}

	//如果 l==r，再移动下
	if l == r {
		l++
		r--
	}

	if left < r {
		QuickSort(left, r, array)
	}
	if right > l {
		QuickSort(l, right, array)
	}
}

func main() {

	var arr [80000]int
	for i:=0; i<80000; i++ {
		rand.Seed(time.Now().UnixNano())
		arr[i] = rand.Intn(900000)
	}

	start := time.Now().Unix()
	QuickSort(0, len(arr)-1, &arr)
	end := time.Now().Unix()
	fmt.Printf("快速排序耗时=%d秒", end-start)

	//arr := [6]int{-9,78,0,23,-567,70}
	//QuickSort(0,5, &arr)
	//fmt.Println(arr)

}
