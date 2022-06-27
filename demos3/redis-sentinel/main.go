package main

import (
	"log"

	"github.com/go-redis/redis"
)

//https://redis.uptrace.dev/guide/go-redis-sentinel.html#redis-server-client
func main() {
	log.Print("hello world.")

	// rdb := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379", // use default Addr
	// 	Password: "",               // no password set
	// 	DB:       0,                // use default DB
	// })

	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "mymaster",
		SentinelAddrs: []string{":26379"},
	})

	//心跳
	pong, err := rdb.Ping().Result()
	log.Println(pong, err) // Output: PONG <nil>
}
