package utility

import (
	"fmt"

	"github.com/go-redis/redis"
)

func ExampleNewClient(host string) {
	client := redis.NewClient(&redis.Options{
		Addr:     host + ":6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println("pong", pong)
	fmt.Println("err", err)
}
