package utility

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// var upgrader = websocket.Upgrader{
// 	CheckOrigin: func(r *http.Request) bool {
// 		return true
// 	},
// }

// func WsHandler(w http.ResponseWriter, r *http.Request) {
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Printf("upgrader error %s\n" + err.Error())
// 		return
// 	}
// 	fmt.Println("conn", conn)
// }

func WsHandler(c chan string) {
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":8000")
	conn, _ := ln.Accept()
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Message Received:", string(message))
	newmessage := strings.ToUpper(message)
	conn.Write([]byte(newmessage + "\n"))
	c <- "launched"
}
