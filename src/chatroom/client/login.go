package main

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"time"
)

//写一个函数，完成登录
func login(userId int, userPwd string) (err error) {

	//下一步就要开始定协议
	//fmt.Printf("userId = %d, userPwd = %s\n", userId, userPwd)
	//
	//return nil

	//1、链接到服务器端
	conn, err := net.Dial("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	//延时关闭
	defer conn.Close()

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
		return
	}

	//5、把data赋给 mes.Data
	mes.Data = string(data)

	//6、将 mes进行序列化
	data, err = json.Marshal(mes.Data)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//7、先把data长度发送给服务器
	//先获取到data的长度->转成一个表示长度的byte切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	//发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail =", err)
		return
	}

	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) fail =", err)
		return
	}

	time.Sleep(5 * time.Second)
	fmt.Println("休眠了5s...")

	//这里还需要处理服务器端返回的消息
	return
}
