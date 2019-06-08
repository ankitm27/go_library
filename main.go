package main

import "Go_library/utility"


func main() {
	colours := []byte(`{"Name":"Eve","Age":6,"Parents":["Alice","Bob"]}`)
	utility.Info(colours)
	// utility.Success(colours)
	// utility.Warning(colours)
	// utility.Error(colours)
}

