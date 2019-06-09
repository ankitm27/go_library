package utility

import (
	"log"
	"net"
	"strconv"
	"strings"
)

func WebClient() {
	ip := "127.0.0.1"
	port := 3333
	message := "Ping"
	StopCharacter := "\r\n\r\n"

	addr := strings.Join([]string{ip, strconv.Itoa(port)}, ":")
	conn, err := net.Dial("tcp", addr)

	defer conn.Close()

	if err != nil {
		log.Fatalln(err)
	}

	conn.Write([]byte(message))
	conn.Write([]byte(StopCharacter))
	log.Printf("Send: %s", message)

	buff := make([]byte, 1024)
	n, _ := conn.Read(buff)
	log.Printf("Receive: %s", buff[:n])

}
