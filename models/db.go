package models

import (
	"github.com/go-redis/redis"
	"os"
)

var client *redis.Client

func Init() {
	client = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("ADDRESS"),
		Password: os.Getenv("PASSWORD"),
	})
}
