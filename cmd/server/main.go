package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sagoresarker/rate-limited-calculator-golang/calculator"
	"github.com/sagoresarker/rate-limited-calculator-golang/config.go"
	"github.com/sagoresarker/rate-limited-calculator-golang/handler"
	"github.com/sagoresarker/rate-limited-calculator-golang/ratelimiter"
	"github.com/sagoresarker/rate-limited-calculator-golang/ratelimiter/storage"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	cfg, err := config.LoadConfig()
	if err != nil {
		e.Logger.Fatal("Error loading config: ", err)
	}

	redisStorage, err := storage.NewRedisStorage(cfg.RedisAddr)
	if err != nil {
		e.Logger.Fatal("Error creating Redis storage: ", err)
	}
	defer redisStorage.Close()

	rateLimiter := ratelimiter.NewRateLimiter(cfg.RateLimitPerMinute, redisStorage)
	calculatorHandler := handler.NewCalculationHandler(calculator.NewRateLimitedCalculator(rateLimiter))

	e.POST("/calculate", calculatorHandler.HandleCalculation)
	e.Logger.Fatal(e.Start(":8080"))
}
