package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"os"
)

var ctx = context.Background() // package level variable?

//Initialize a new Redis client

func newRedisClient() *redis.Client {
	var addr = os.Getenv("REDIS")
	if addr == "" {
		addr = "localhost"
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr + ":6379",
		Password: "",
		DB:       0,
	})
	return rdb
}
