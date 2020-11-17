package main

import (
	"fmt"
	"time"
)

//函数
func sayHello()  {
	for i:=0; i<10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello, world")
	}
}

//函数
func test()  {
	//这里使用错误处理机制 defer + recover
	defer func() {
		//捕获test抛出的panic
		if err := recover(); err != nil {
			fmt.Println("test() 发生错误", err)
		}
	}()
	//定义一个map
	var myMap map[int]string
	myMap[0] = "golang"
}

func main() {
	go sayHello()
	go test()

	for i:=0; i<10; i++ {
		fmt.Println("main() ok=", i)
		time.Sleep(time.Second)
	}
}
