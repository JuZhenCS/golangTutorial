//这只是一个实验程序,所以很多错误判断都没写.
//在实用环境中一定要!
//golang的核心优势之一就是,只要进行了正确的判断,就没有段错误.
package main

import (
	"encoding/json"
	"net"
)

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:1234")
	var myBlog Blog    //定义结构体
	myBlog.Name = "赵昊" //赋值
	myBlog.Url = "www.zhaoritian.com"
	buffer, _ := json.Marshal(myBlog) //编码
	conn.Write(buffer)
	return
}

type Blog struct {
	Name string
	Url  string
}
