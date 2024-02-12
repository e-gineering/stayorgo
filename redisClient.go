package main

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background() // package level variable?

//Initialize a new Redis client

func newRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return rdb
}
