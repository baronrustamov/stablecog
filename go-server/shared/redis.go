package shared

import (
	"log"

	"github.com/go-redis/redis/v8"
)

var Redis *redis.Client

func SetupRedis() {
	opts, err := redis.ParseURL(GetEnv("REDIS_CONNECTION_STRING"))
	if err != nil {
		log.Printf("Error parsing REDIS_CONNECTION_STRING: %v", err)
	}
	Redis = redis.NewClient(opts)
	res, err := Redis.Ping(ctx).Result()
	if err != nil {
		log.Printf("Error pinging Redis: %v", err)
	}
	log.Printf("Redis ping result: %v", res)
}
