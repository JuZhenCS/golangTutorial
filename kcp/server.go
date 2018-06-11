package main

import (
	"fmt"
	"github.com/xtaci/kcp-go"
)

func main() {
	listener, err := kcp.Listen("localhost:1234") //代替net建立连接
	if err != nil {                               //极度冗余的错误判断，因为网络是不可靠的，可能在任何时间出问题
		fmt.Println(err.Error())
		return
	}
	defer listener.Close()
	conn, err := listener.Accept() //接到客户端连接
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()
	buffer := make([]byte, 1024)
	length, err := conn.Read(buffer) //读取客户端消息
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(buffer[:length])) //打印消息
}
