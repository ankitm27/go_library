package utility

import (
	"bufio"
	"fmt"
	"net"
)

func WebClient() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8080")
	var a = "check"
	fmt.Fprintf(conn, a+"\n")
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println("message", message)
}
