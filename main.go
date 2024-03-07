package main

// Roseville, CA, US Coords: 38.76, -121.33

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/redis/go-redis/v9"
)

func main() {

	// connect to redis
	ctx := context.Background()
	rclient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer rclient.Close()

	// init reader for user input
	reader := bufio.NewReader(os.Stdin)

	// forecast region hardcoded for Roseville, CA
	region := "STO/48,75"

	for {
		fmt.Println("Enter an option:")
		fmt.Println("1. Get weather from api.weather.gov")
		fmt.Println("2. Get cached forecast from redis")

		// get user input and trim whitespace
		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		// take action on option selected
		if option == "1" {
			fmt.Println("Getting weather from api.weather.gov")

			// get weather forecast
			forecastBytes, err := getForecast(region)
			if err != nil {
				log.Fatal(err)
			}

			// convert bytes to usable objects
			data, err := unmarshalForecast(forecastBytes)
			if err != nil {
				log.Fatal(err)
			}

			// cache forecast in redis
			err = storeForecastRedis(rclient, ctx, region, data)
			if err != nil {
				log.Fatal(err)
			}

			// display data received
			displayForecastPeriods(data.Properties.Periods)

			break
		} else if option == "2" {
			fmt.Println("Getting cached weather from redis")

			// get data from redis
			data, err := fetchForecastRedis(rclient, ctx, region)
			if err != nil {
				log.Fatal(err)
			}

			// display data received
			displayForecastPeriods(data.Properties.Periods)

			break
		} else {
			fmt.Println("Invalid option.")
		}
	}

}
