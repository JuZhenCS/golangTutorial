package main

import (
	"net"
	"time"
)

func main() {
	conn := getConn() //建立连接
	for {
		_, err := conn.Write([]byte("Hello World!")) //向服务端发数据
		if err != nil {                              //发送数据出错就重新建立连接
			conn = getConn()
		}
		time.Sleep(10 * time.Second) //睡眠10秒
	}
}

func getConn() net.Conn {
	for {
		conn, err := net.Dial("tcp", "23.106.153.177:1234") //建立连接
		if err != nil {                                     //出错就重新建立连接
			time.Sleep(1 * time.Second)
			continue
		}
		return conn
	}

}
