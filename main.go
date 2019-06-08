package main

import "Go_library/utility"


func main() {
	colours := map[string]interface{}{
		"check":map[string]interface{}{
			"check":"check",
			"check1":1,
		},
	}
	utility.Info(colours)
	utility.Success(colours)
	utility.Warning(colours)
	utility.Error(colours)
}

