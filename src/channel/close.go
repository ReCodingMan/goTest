package main

import "fmt"

func main() {
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	close(intChan)

	//intChan <- 300
	fmt.Printf("ok~")

	//遍历管道
	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i*2
	}
	close(intChan2)

	for v := range intChan2 {
		fmt.Printf("v=%v \n", v)
	}
}
