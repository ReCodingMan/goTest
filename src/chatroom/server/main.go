package main

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {

	buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据...")
	_, err = conn.Read(buf[:4])
	if err != nil {
		//err = errors.New("conn.Read err")
		return
	}

	//根据buf[:4] 转成一个 uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])

	//根据 pkgLen 读取消息内容
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		//err = errors.New("conn.Read fail err")
		return
	}

	//把 pkgLen 反序列化成 -> message.Message
	//技术就是一层窗户纸 &mes
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	return
}

//编写专门处理登录的请求方法
func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
	//核心代码
	//1、先从mes中取出 mes.Data，直接反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}

	//先声明一个 resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	//如果用户id=100，密码=123456，认为合法，否则不合法
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		//合法
	} else {
		//不合法
	}
	return
}

//编写一个ServerProcessMes 函数
//功能：根据客户端发送消息种类不同，决定调用哪个函数来处理
func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {

	switch mes.Type {
		case message.LoginMesType:
			//处理登录逻辑
			err = serverProcessLogin(conn, mes)
		case message.RegisterMesType:
			//处理注册逻辑
		default:
			fmt.Println("消息类型不存在，无法处理...")
	}
	return
}

//处理和客户端的通讯
func process(conn net.Conn) {

	//延时关闭
	defer conn.Close()

	//循环读客户端发送的信息
	for {

		//这里，我们将读取数据包，直接封装成一个方法，返回 message，err
		mes, err := readPkg(conn)

		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务端也退出")
				return
			} else {
				fmt.Println("readPkg err=", err)
			}
		}
		fmt.Println("mes=", mes)
	}

}

func main()  {

	//提示信息
	fmt.Println("服务器在8889监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
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
