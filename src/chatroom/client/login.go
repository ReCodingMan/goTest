package main

import (
	"chatroom/common/message"
	"encoding/json"
	"fmt"
	"net"
)

//写一个函数，完成登录
func login(userId int, userPwd string) (err error) {

	//下一步就要开始定协议
	//fmt.Printf("userId = %d, userPwd = %s\n", userId, userPwd)
	//
	//return nil

	//1、链接到服务器端
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}

	//2、准备通过conn发送消息给服务
	var mes message.Message
	mes.Type = message.LoginMesType
	//3、创建一个 loginMes 结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//4、将 loginMes 序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("loginMes Marshal err=", err)
	}

	//5、把data赋给 mes.Data
	mes.Data = string(data)

	//6、将 mes进行序列化
	data, err = json.Marshal(mes.Data)
	if err != nil {
		fmt.Println("mes.Data Marshal err=", err)
	}

	//7、先把data长度发送给服务器
	//先获取到data的长度->转成一个表示长度的byte切片


}
