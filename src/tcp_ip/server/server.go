package main

import (
	"fmt"
	"net" //做网络socket开发时，net 包含我们需要所有的方法和函数
)

func process(conn net.Conn) {

	defer conn.Close() // 关闭conn

	for {
		//创建一个新的切片
		buf := make([]byte, 1024)
		//这里循环接收客户端发送的数据

		//conn.Read(buf)
		//1、等待客户端通过 conn 发送信息
		//2、如果客户端没有write[发送]，那么协程就阻塞在这里
		fmt.Printf("服务器在等待客户端%s，发送信息\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("服务器端read出错", err)
			return
		}

		//3、显示客户端发送的数据内容到服务器的终端
		fmt.Print(string(buf[:n]))
	}
}

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
		fmt.Printf("Accept() suc con=%v, 客户端ip=%v \n", conn, conn.RemoteAddr().String())

		//这里起一个协程，为客户端服务
		go process(conn)
	}
	fmt.Printf("listen=%v\n", listen)

}
