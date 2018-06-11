//这只是一个实验程序,所以很多错误判断都没写.
//在实用环境中一定要加上!
//golang的核心优势之一就是,只要对错误进行了正确的判断,就没有段错误.
package main

import (
	"encoding/json"
	"fmt"
	"net"
)

func main() {
	listener, _ := net.Listen("tcp", "localhost:1234") //_表示空变量,赋值但不使用
	conn, _ := listener.Accept()
	buffer := make([]byte, 1024)
	length, _ := conn.Read(buffer)
	var myBlog Blog                          //定义接收的数据结构
	json.Unmarshal(buffer[:length], &myBlog) //解码,结构体传址
	fmt.Println(myBlog.Name, myBlog.Url)
}

type Blog struct { //json的结构,注意,全部名字都要首字母大写
	Name string
	Url  string
}
