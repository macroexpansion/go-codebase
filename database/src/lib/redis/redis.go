package redis

import (
	"context"
	"github.com/go-redis/redis/v9"
)

func Connect(address string) (*redis.Client, error) {
	ctx := context.TODO()
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return client, nil
}
