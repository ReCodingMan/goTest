package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

//定义一个全局的pool
var pool *redis.Pool

//当程序启动时，就初始化连接池
func init()  {
	pool = &redis.Pool{
		MaxIdle: 8,	//最大空闲链接数量
		MaxActive: 0,	//表示和数据库的最大连接数，0表示没有限制
		IdleTimeout: 100,	//最大空闲时间
		Dial: func() (redis.Conn, error) {	//初始连接的代码，连接哪个ip
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {

	//先从pool取出一个链接
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("Set", "name", "汤姆猫～")
	if err != nil {
		fmt.Println("conn Do err=", err)
		return
	}

	//取
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("conn Get err=", err)
		return
	}
	fmt.Println("r=", r)
}