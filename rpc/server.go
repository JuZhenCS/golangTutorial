package main

import (
	"net"
	"net/rpc"
)

func main() {
	myProgram := new(Program)
	rpc.Register(myProgram) //注册rpc
	listener, _ := net.Listen("tcp", ":1234")
	rpc.Accept(listener) //监听rpc
	listener.Accept()
}

type Program int //rpc的接口

func (t Program) Add(number int, result *int) error { //rpc的方法
	*result = number + 1 //输入函数随意,输出函数用指针
	return nil
}
