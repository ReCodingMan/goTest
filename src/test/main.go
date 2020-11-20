package main

import "fmt"

func main() {
	var sli []int
	var result []int
	len := 5
	sli = make([]int, 6)
	result = make([]int, 0)
	sli = []int{1,2,3,1,2,3}

	for i:=0; i<=6-len; i++ {
		result = append(result, maxSli(sli[i:i+3]))
	}
	fmt.Println(result)
}

func maxSli(sli []int) int {
	maxV := sli[0]
	for _,v := range sli {
		if v > maxV {
			maxV = v
		}
	}
	return maxV
}