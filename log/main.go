package main

import (
	"log"
)

func main() {
	log.SetPrefix("日志输出的例子:")                       //日志前缀
	log.SetFlags(log.LstdFlags | log.Lmicroseconds) //日志格式
	//打印错误日志,黄色字体
	log.Printf("%c[33m %s %c[0m", 0x1B, "this is the error log!", 0x1B)
	//打印异常日志,红色字体
	log.Fatalf("%c[31m %s %c[0m", 0x1B, "this is the panic log!", 0x1B)
}
