package main

import (
	"fmt"
	"net"
	"net/rpc"
)

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:1234") //连接服务器
	c := rpc.NewClient(conn)                     //连接rpc
	var result int
	c.Call("Program.Add", 1, &result) //远程函数调用
	fmt.Println(result)
}
