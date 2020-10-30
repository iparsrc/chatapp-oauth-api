package domain

import (
	"log"
	"os"

	"github.com/go-redis/redis"
)

var (
	client *redis.Client
)

func init() {
	connectToRedis()
}

func connectToRedis() {
	dsn := os.Getenv("REDIS_DSN")
	if dsn == "" {
		dsn = "localhost:6379"
	}
	client = redis.NewClient(&redis.Options{
		Addr: dsn,
	})
	if _, err := client.Ping().Result(); err != nil {
		log.Fatal("error: can't connect to redis.")
	}
}
