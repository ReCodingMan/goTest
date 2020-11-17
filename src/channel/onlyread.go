package main

import "fmt"

func main() {
	//管道可读可写
	//var chan1 chan int

	//管道只写
	var chan2 chan<- int
	chan2 = make(chan int, 3)
	chan2 <- 1
	chan2 <- 2
	chan2 <- 3
	//num := <-chan2 //错误，只读类型
	fmt.Println("chan2 =", chan2)
	//fmt.Println("num =", num)

	//管道只读
	var chan3 <-chan int
	num2 := <-chan3
	fmt.Println("num2 =", num2)
}
