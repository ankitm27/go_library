package utility

import "fmt"

func Info(data map[string]interface{})  {
	fmt.Println(data["check"])
}

func Success(data map[string]interface{}){
	fmt.Println(data)
}

func Warning(data map[string]interface{}){
	fmt.Println(data)
}

func Error(data map[string]interface{}){
	fmt.Println(data)
}

