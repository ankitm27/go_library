package utility

import (
	"encoding/json"
	"fmt"
	"log"
)

func printJson(data map[string]interface{}) {
	b, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}

func Info(data map[string]interface{}) {
	printJson(data)
}

func Success(data map[string]interface{}) {
	printJson(data)
}

func Warning(data map[string]interface{}) {
	fmt.Println(data)
}

func Error(data map[string]interface{}) {
	printJson(data)
}
