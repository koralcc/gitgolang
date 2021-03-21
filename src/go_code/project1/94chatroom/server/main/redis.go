package main

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

//定义一个全局pool
var pool *redis.Pool

func initPool(address string, maxIdle, maxActive int, idleTimeout time.Duration) {
	pool = &redis.Pool{
		MaxIdle:     maxIdle,     //最大空闲连接数
		MaxActive:   maxActive,   //表示和数据库的最大链接数，0表示没有限制
		IdleTimeout: idleTimeout, //最大空闲时间
		Dial: func() (redis.Conn, error) { //初始化连接，链接哪个ip的redis
			return redis.Dial("tcp", address)
		},
	}
}
