package utility

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

var isConnectionAvailable = false
var Conn net.Conn
var err error

func WebClient(message string) {
	StopCharacter := "\r\n\r\n"
	// addr := strings.Join([]string{BackendUrl, strconv.Itoa(BackendPort)}, ":")
	// conn, err := net.Dial("tcp", addr)
	// fmt.Println(reflect.TypeOf(conn))
	// defer conn.Close()
	// if err != nil {
	// log.Fatalln(err)
	// }

	// if !isConnectionAva == "" {
	// 	WebClientConnection()
	// }
	Conn := WebClientConnection()
	fmt.Println("conn", Conn)
	Conn.Write([]byte(message))
	Conn.Write([]byte(StopCharacter))
	log.Printf("Send: %s", message)
	buff := make([]byte, 1024)
	n, _ := Conn.Read(buff)
	log.Printf("Receive: %s", buff[:n])
}

func WebClientConnection() net.Conn {
	if !isConnectionAvailable {
		addr := strings.Join([]string{BackendUrl, strconv.Itoa(BackendPort)}, ":")
		Conn, err = net.Dial("tcp", addr)
		defer Conn.Close()
		if err != nil {
			log.Fatalln(err)
		}
		isConnectionAvailable = true
		// return Conn
	}
	return Conn
}
