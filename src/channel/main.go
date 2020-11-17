package main

import "fmt"

type Cat struct {
	Name string
	Age int
}

func main()  {

	var allChan chan interface{}
	allChan = make(chan interface{}, 3)

	allChan <- 10
	allChan <- "tom cat"
	cat := Cat{"小花猫",4}
	allChan <- cat

	//直接取第三个
	<- allChan
	<- allChan
	newCat := <- allChan

	fmt.Printf("newCat=%T, newCat=%v \n", newCat, newCat)
	a := newCat.(Cat)//类型断言
	fmt.Printf("newCat.Name=%v", a.Name)
}
