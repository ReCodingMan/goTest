package main

import "fmt"

func writeData(intChan chan int)  {
	for i:=1; i<=50; i++ {
		//放入数据
		intChan <- i
		fmt.Printf("writeData 写入数据=%v\n", i)
	}
	close(intChan)
}

func readData(intChan chan int, boolChan chan bool) {
	for {
		v, ok := <- intChan
		if !ok {
			break
		}
		fmt.Printf("readData 读到数据=%v\n", v)
	}
	//任务完成
	boolChan <- true
	close(boolChan)
}

func main()  {

	// 创建两个管道
	intChan := make(chan int, 50)
	boolChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, boolChan)

	for {
		_, ok := <- boolChan
		if !ok {
			break
		}
	}
}
