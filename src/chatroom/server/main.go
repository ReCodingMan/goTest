package main

import (
	"fmt"
	"net"
)

//处理和客户端的通讯
func process(conn net.Conn) {

	//读客户端发送的信息

}

func main()  {

	//提示信息
	fmt.Println("服务器在8889监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}

	//一旦监听成功，就等待客户端来链接服务器
	for {
		fmt.Println("等待客户端来链接服务器。。。")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept err=", err)
		}

		//一旦链接成功，则启动一个协程和客户端保持通讯...
		go process(conn)
	}
}
