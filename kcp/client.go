package main

import (
	"fmt"
	"github.com/xtaci/kcp-go"
)

func main() {
	conn, err := kcp.Dial("localhost:1234")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()
	conn.Write([]byte("Hello World!"))
	return
}
