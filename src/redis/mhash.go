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
	_, err = conn.Do("HMSet", "user02", "name", "mali", "age", "18")
	if err != nil {
		fmt.Println("set err=", err)
		return
	}

	//3、通过go向redis读取数据 string
	r, err := redis.Strings(conn.Do("HMGet", "user02", "name", "age"))
	if err != nil {
		fmt.Println("HMGet err=", err)
		return
	}
	fmt.Printf("HMGet suc=%v\n", r)

	for i, v := range r {
		fmt.Printf("i=%v, v=%v\n", i, v)
	}
}
