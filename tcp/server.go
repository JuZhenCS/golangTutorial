package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":1234") //声明连接
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer listener.Close() //默认关闭连接
	for {                  //循环接受连接
		conn, err := listener.Accept() //接受连接
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		go doReceiver(conn) //收到连接后多放到协程处理
	}
}

func doReceiver(conn net.Conn) {
	defer conn.Close() //默认关闭连接
	buffer := make([]byte, 1024)
	for {
		length, err := conn.Read(buffer) //接受数据
		remoteaddr := conn.RemoteAddr().String()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("接收到数据：", string(buffer[:length])) //打印数据
		fmt.Println(remoteaddr)
	}
}
