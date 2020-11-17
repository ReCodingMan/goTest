package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now().Unix()
	for i:=0; i<200000; i++ {
		flag := true
		//判断这个数是不是素数
		for j:=2; j<i; j++ {
			if i % j == 0 {
				//说明该数不是素数
				flag = false
				break
			}
		}
		if flag {
			//是素数
			//fmt.Println("是的")
		}
	}

	end := time.Now().Unix()
	fmt.Println("使用时间耗时为=", end - start)

}