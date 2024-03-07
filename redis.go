package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

// stores/caches given APIData into redis at the given key (region) with a
// hardcoded 10 second expiration for demonstration purposes
func storeForecastRedis(rclient *redis.Client, ctx context.Context, key string, data *APIData) error {
	// serialize APIData to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// store JSON data in redis with hardcoded 10 second expiration
	err = rclient.Set(ctx, key, jsonData, 10*time.Second).Err()
	if err != nil {
		return err
	}

	return nil
}

// gets the cached forecast APIData from redis (if exists) and returns a pointer
// to an APIData struct. errors if the cached forecast is not in redis
func fetchForecastRedis(rclient *redis.Client, ctx context.Context, key string) (*APIData, error) {
	// get JSON data from redis
	data, err := rclient.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	// deserialize JSON -> APIData struct
	var apidata APIData
	err = json.Unmarshal([]byte(data), &apidata)
	if err != nil {
		return nil, err
	}

	return &apidata, nil
}
