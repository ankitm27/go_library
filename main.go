package main

import (
	"go_library/utility"
)

var serverAddress = ":8080"

func main() {
	// colours := []byte(`{"Name":"Eve","Age":6,"Parents":["Alice","Bob"]}`)
	// utility.Info(colours)
	// utility.Success(colours)
	// utility.Warning(colours)
	// utility.Error(colours)
	// const host = utility.BaseUrlRedis
	// fmt.Println(host)
	// utility.ExampleNewClient(host)
	// http.HandleFunc("/ws", utility.WsHandler)
	// utility.Handler()
	// log.Fatal(http.ListenAndServe(serverAddress, nil))
	// const host = utility.BackendUrl
	// const port = utility.BackendPort
	// fmt.Println("host", host)
	// fmt.Println("port", port)
	// utility.WebClientConnection()
	b := []byte(`{"Name":"Eve","Age":6,"Parents":["Alice","Bob"]}`)
	var message = string(b)
	utility.Info(message)
	// utility.WebClient(host, port, message)

}
