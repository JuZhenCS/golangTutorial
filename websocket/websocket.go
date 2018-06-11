package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{ //websocket接口
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Message struct { //json对象,首字母大写
	Name   string
	Localx float64
	Localy float64
}

var message = make(chan Message)             //要发送的数据
var clients = make(map[*websocket.Conn]bool) //连接数组

func main() {
	go sendMessage()                            //发消息
	go http.HandleFunc("/websocket", webSocket) //启动websocket
	go http.ListenAndServe(":3000", nil)        //启动服务
	getInput()                                  //用户输入
}

func webSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil) //setup websocket
	chencError(err)
	fmt.Println(conn.RemoteAddr().String())
	clients[conn] = true //push conn to list
	var msg Message
	for {
		err = conn.ReadJSON(&msg)
		chencError(err)
		message <- msg
		fmt.Println(msg.Name)
	}
}

func sendMessage() {
	for {
		msg := <-message
		fmt.Println(msg)
		for client := range clients {
			err := client.WriteJSON(msg)
			chencError(err)
			if err != nil {
				fmt.Println("Send error")
				delete(clients, client)
				client.Close()
			}
		}
	}
}

func getInput() {
	var name string
	var direction string
	var msg Message
	for {
		fmt.Println("input name and direction")
		fmt.Scanln(&name, &direction)
		switch name {
		case "1":
			msg.Name = "feiji001"
		case "2":
			msg.Name = "feiji002"
		default:
			fmt.Println("input name as 1 or 2")
			continue
		}
		switch direction {
		case "h":
			msg.Localx = -0.001
			msg.Localy = 0
		case "j":
			msg.Localx = 0
			msg.Localy = -0.001
		case "k":
			msg.Localx = 0
			msg.Localy = 0.001
		case "l":
			msg.Localx = 0.001
			msg.Localy = 0
		default:
			fmt.Println("input direction as hjkl")
			continue
		}
		message <- msg
	}
}

func chencError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
