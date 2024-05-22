package main

import (
	"fmt"

	"github.com/sagoresarker/rate-limited-calculator-golang/calculator"
	"github.com/sagoresarker/rate-limited-calculator-golang/config.go"
	"github.com/sagoresarker/rate-limited-calculator-golang/ratelimiter"
	"github.com/sagoresarker/rate-limited-calculator-golang/ratelimiter/storage"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("error loading config: %v\n", err)
		return
	}

	redisStorage, err := storage.NewRedisStorage(cfg.RedisAddr)
	if err != nil {
		fmt.Printf("error creating redis storage: %v\n", err)
		return
	}

	defer redisStorage.Close()

	rateLimiter := ratelimiter.NewRateLimiter(cfg.RateLimitPerMinute, redisStorage)
	calculator := calculator.NewRateLimitedCalculator(rateLimiter)

	for i := 0; i < 20; i++ {
		result, err := calculator.Add(1, 2)
		if err != nil {
			fmt.Printf("error adding numbers: %v\n", err)
		}

		fmt.Printf("result: %d\n", result)
	}

	fmt.Println("done")

	for i := 0; i < 20; i++ {
		result, err := calculator.Subtract(1, 2)
		if err != nil {
			fmt.Printf("error adding numbers: %v\n", err)
			return
		}

		fmt.Printf("result: %d\n", result)
	}
}
