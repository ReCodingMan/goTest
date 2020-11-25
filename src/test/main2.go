package main

import "fmt"

func smallerNumbersThanCurrent(nums []int) []int {
	arr := make([]int, 0101)
	for i := 0; i < len(nums); i++ {
		arr[nums[i]]++
	}
	fmt.Println(arr)

	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			countBefore := count
			count += arr[i]
			arr[i] = countBefore
		}
	}
	fmt.Println(arr)

	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = arr[nums[i]]
	}
	return res
}

func main()  {
	arr := []int {8,1,2,2,3}
	fmt.Println(smallerNumbersThanCurrent(arr))
}
