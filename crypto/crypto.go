//这是一个账号密码加密解密的例子
//程序员开发程序的基本礼貌,如果数据被泄露了,密码也不应该能被还原
//date:2018-06-12
//此时的感想:不想加班,我喜欢写文档,但是不想做需要用office的工作
//需要安装bcyrpt库
//go get golang.org/x/crypto/bcrypt

package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	passwddok := "admin"                                                            //正确的密码
	passwdERR := "adminerr"                                                         //错误的密码
	hash, err := bcrypt.GenerateFromPassword([]byte(passwddok), bcrypt.DefaultCost) //编码
	checkErr(err)
	encodePW := string(hash) //编码后的密码,每次生成都是不同的,但是保存任意一份即可

	err = bcrypt.CompareHashAndPassword([]byte(encodePW), []byte(passwddok))
	checkErr(err)
	fmt.Println("密码正确")
	err = bcrypt.CompareHashAndPassword([]byte(encodePW), []byte(passwdERR))
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
