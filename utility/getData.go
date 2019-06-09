package utility

import (
	"encoding/json"
	"fmt"
)

func printJson(dataObj []byte) {
	var v interface{}
	json.Unmarshal(dataObj, &v)
	data := v.(map[string]interface{})

	for k, v := range data {
		switch v := v.(type) {
		case []interface{}:
			fmt.Println(k)
			for i, u := range v {
				fmt.Println("    ", i, u)
			}
		default:
			fmt.Println(k, v)
		}
	}
}

func Info(data string) {
	WebClient(data)
}

func Success(data []byte) {
	printJson(data)
}

func Warning(data []byte) {
	fmt.Println(data)
}

func Error(data []byte) {
	printJson(data)
}
