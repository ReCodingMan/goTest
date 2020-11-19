package main

import (
	"fmt"
	"net" //做网络socket开发时，net 包含我们需要所有的方法和函数
)

func main()  {

	fmt.Println("服务器开始监听...")
	//网络协议是tcp
	//监听端口是8888
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err = ", err)
	}

	defer listen.Close() //延时关闭

	for {
		//等待客户端来连接
		fmt.Println("等待客户端来连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err = ", err)
		}
		fmt.Printf("Accept() suc con=%v\n", conn)

		//这里起一个协程，为客户端服务
	}
	fmt.Printf("listen=%v\n", listen)

}
