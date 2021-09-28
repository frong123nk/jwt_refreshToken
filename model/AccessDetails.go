package model

import (
	"github.com/go-redis/redis/v8"
)

type AccessDetails struct {
	AccessUuid string
	UserId     uint64
}

func Rclient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client
}
