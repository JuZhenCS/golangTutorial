package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"sync"
	"time"
)

var myRedisPool *redis.Pool //redis连接池

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     100,             //最大空闲连接数量
		MaxActive:   1000,            //最大同时存在连接数量
		IdleTimeout: 3 * time.Second, //超过最大空闲连接数量的连接，3秒不用就销毁
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				"tcp",
				"localhost:6379",
				//redis.DialConnectTimeout(500*time.Millisecond), //连接超时
				//redis.DialReadTimeout(500*time.Millisecond), //读取超时
				//redis.DialWriteTimeout(500*time.Millisecond),   //写入超时
			)
		},
	}
}

func init() {
	myRedisPool = newPool() //初始化连接
}

var waitGroup sync.WaitGroup //同步信号

func main() {
	waitGroup.Add(1) //同步信号+1
	go listen()
	waitGroup.Wait()
}

func listen() {
	client := myRedisPool.Get()
	defer client.Close()
	ScribClient := redis.PubSubConn{client}
	ScribClient.Subscribe("redisChat")
	for {
		switch v := ScribClient.Receive().(type) {
		case redis.Message:
			fmt.Println(v.Channel, ": ", string(v.Data))
		case redis.Subscription:
			fmt.Println(v.Channel, " ", v.Kind, " ", v.Count)
		case error:
			fmt.Println(v)
			return
		}
	}
	waitGroup.Done()
}
