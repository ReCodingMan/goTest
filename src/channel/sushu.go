package main

import (
	"fmt"
	"time"
)

//初始化管道
func PutNum(intChan chan int) {
	for i:=0; i<200000; i++ {
		intChan <- i
	}
	close(intChan)
}

//从 intChan 取出数据，并判断是否为素数
func PrimeNum(intChan chan int, primeChan chan int, resultChan chan bool)  {

	//使用 for 循环
	var flag bool
	for  {
		num,ok := <-intChan
		if !ok {
			//管道取不到
			break
		}
		flag = true
		//判断这个数是不是素数
		for i:=2; i<num; i++ {
			if num % i == 0 {
				//说明该数不是素数
				flag = false
				break
			}
		}
		if flag {
			//是素数
			primeChan <- num
		}
	}

	fmt.Println("有一个协程取不到数据，退出")
	//这里不能关闭管道
	resultChan <- true
}

func main() {

	intChan := make(chan int, 8000)
	primeChan := make(chan int, 200000) //放入结果
	//标识退出的管道
	resultChan := make(chan bool, 4)

	start := time.Now().Unix()
	//开启一个线程，向 intChan 放入1-8000个数
	go PutNum(intChan)

	//开启4个线程，从 intChan 取出数据，并判断是否为素数
	for i:=0; i<8; i++ {
		go PrimeNum(intChan, primeChan, resultChan)
	}

	go func() {
		for i:=0; i<4; i++ {
			<-resultChan
		}

		end := time.Now().Unix()
		fmt.Println("使用时间耗时为=", end - start)

		close(primeChan)
	}()

	for {
		_, ok := <-primeChan
		if !ok {
			break
		}
		//fmt.Printf("素数=%d\n", res)
	}
	fmt.Println("main线程退出")
}