package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {

	//1、连接redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}
	defer conn.Close()

	//2、通过go向redis写入数据 string [key-val]
	_, err = conn.Do("HSet", "user01", "name", "cheng")
	if err != nil {
		fmt.Println("set err=", err)
		return
	}

	_, err = conn.Do("HSet", "user01", "age", "20")
	if err != nil {
		fmt.Println("set err=", err)
		return
	}

	//3、通过go向redis读取数据 string
	r1, err := redis.String(conn.Do("HGet", "user01", "name"))
	if err != nil {
		fmt.Println("get err=", err)
		return
	}
	fmt.Println("get suc", r1)

	r2, err := redis.String(conn.Do("HGet", "user01", "age"))
	if err != nil {
		fmt.Println("get err=", err)
		return
	}
	fmt.Println("get suc", r2)
}
