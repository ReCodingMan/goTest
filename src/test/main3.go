package main

import (
	"fmt"
	"strings"
)

func numJewelsInStones(J string, S string) int {
	res := 0
	for i:=0; i<len(J); i++ {
		res += strings.Count(S, J[i:i+1])
	}
	return res
}

func main() {
	J := "aA"
	S := "aAAbbbb"
	fmt.Println(numJewelsInStones(J, S))
}
