package models

import "github.com/go-redis/redis"

var client *redis.Client

func Init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "redis-17838.c92.us-east-1-3.ec2.cloud.redislabs.com:17838",
		Password: "5ifhOBhLO8FCMKae7pQ7N0BKtjRU626a",
	})
}
