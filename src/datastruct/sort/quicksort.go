package main

import "fmt"

func QuickSort(left int, right int, array *[6]int) {
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

	arr := [6]int{-9,78,0,23,-567,70}
	QuickSort(0,5, &arr)
	fmt.Println(arr)

}
