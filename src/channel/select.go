package main

import "fmt"

func main()  {

	//定义一个管道 10个数字int
	intChan := make(chan int, 10)
	for i:=0; i<10; i++ {
		intChan<-i
	}

	stringChan := make(chan string, 5)
	for i:=0; i<5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	label:
	for {
		select {
			case v := <-intChan:
				//注意：这里，如果intChan一直没有关闭，不会一直阻塞而 deadlock
				//会自动到下一个case匹配
				fmt.Printf("从 intChan 读取的数据%d\n", v)
			case v := <-stringChan:
				fmt.Printf("从 stringChan 读取的数据%s\n", v)
			default:
				fmt.Println("都取不到")
				break label
		}
	}
}
