package main

import "Go_library/utility"

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

	utility.WebClient()

}
