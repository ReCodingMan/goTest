package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	conn, err := net.Dial("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}
	fmt.Println("conn 成功=", conn)

	//功能一：输出单行数据，然后退出
	reader := bufio.NewReader(os.Stdin) //标准输入

	for {
		//从终端读取一行用户输入，准备发送服务器
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("ReadString err=", err)
		}
		line = strings.Trim(line, " \r\n")
		if line == "exit" {
			fmt.Println("客户端退出")
			break
		}

		//发送给服务器
		n, err := conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("conn.Write err=", err)
		}
		fmt.Printf("客户端发送了 %d 字节的数据，并退出 \n", n)
	}

}