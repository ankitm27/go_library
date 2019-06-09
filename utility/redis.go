package utility

import (
	"fmt"

	"github.com/go-redis/redis"
)

func RedisClient(host string) {
	client := redis.NewClient(&redis.Options{
		Addr:     host + ":6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("pong", pong)
}
