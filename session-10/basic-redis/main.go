package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	// Create a Redis client
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Set a key-value pair
	ctx := context.Background()
	err := client.Set(ctx, "exampleKey", "exampleValue", 0).Err()
	if err != nil {
		panic(err)
	}

	// Set a key-value pair
	err = client.Set(ctx, "exampleKey", "exampleValue", 0).Err()
	if err != nil {
		panic(err)
	}

	// Get the value by key
	value, err := client.Get(ctx, "exampleKey").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Value:", value)
}
