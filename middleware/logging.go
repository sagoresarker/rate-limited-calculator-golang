package middleware

import (
	"log"

	"github.com/labstack/echo/v4"
)

func LoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Printf("Received request: %v %v", c.Request().Method, c.Request().URL)
		return next(c)
	}
}
