package main

import (
	"Go_library/utility"
	"fmt"
)

func main() {
	// colours := []byte(`{"Name":"Eve","Age":6,"Parents":["Alice","Bob"]}`)
	// utility.Info(colours)
	// utility.Success(colours)
	// utility.Warning(colours)
	// utility.Error(colours)
	const host = utility.BaseUrlRedis
	fmt.Println(host)
	// utility.ExampleNewClient(host)
	// http.HandleFunc("/ws", utility.Handler)
	// utility.Handler()
}
