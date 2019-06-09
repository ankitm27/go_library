package main

import (
	"go_library/utility"
)

var serverAddress = ":8080"

func main() {
	// colours := []byte(`{"Name":"Eve","Age":6,"Parents":["Alice","Bob"]}`)
	b := []byte(`{"Name":"Eve","Age":6,"Parents":["Alice","Bob"]}`)
	var message = string(b)
	utility.Info(message)
}
