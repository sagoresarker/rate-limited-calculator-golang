package calculator

import "github.com/sagoresarker/rate-limited-calculator-golang/ratelimiter"

type RateLimitedCalculator struct {
	rateLimiter ratelimiter.RateLimiter
}

func NewRateLimitedCalculator(rateLimiter ratelimiter.RateLimiter) *RateLimitedCalculator {
	return &RateLimitedCalculator{
		rateLimiter: rateLimiter,
	}
}

func (rlc *RateLimitedCalculator) Add(a, b int) (int, error) {
	allowed, err := rlc.rateLimiter.AllowRequest()

	if err != nil {
		return 0, err
	}

	if !allowed {
		return 0, err
	}

	return a + b, nil
}

func (rlc *RateLimitedCalculator) Subtract(a, b int) (int, error) {
	allowed, err := rlc.rateLimiter.AllowRequest()

	if err != nil {
		return 0, err
	}

	if !allowed {
		return 0, err
	}

	return a - b, nil
}

func (rlc *RateLimitedCalculator) Multiply(a, b int) (int, error) {
	allowed, err := rlc.rateLimiter.AllowRequest()

	if err != nil {
		return 0, err
	}

	if !allowed {
		return 0, err
	}

	return a * b, nil
}

func (rlc *RateLimitedCalculator) Divide(a, b int) (int, error) {
	allowed, err := rlc.rateLimiter.AllowRequest()

	if err != nil {
		return 0, err
	}

	if !allowed {
		return 0, err
	}

	return a / b, nil
}

func (rlc *RateLimitedCalculator) Modulo(a, b int) (int, error) {
	allowed, err := rlc.rateLimiter.AllowRequest()

	if err != nil {
		return 0, err
	}

	if !allowed {
		return 0, err
	}

	return a % b, nil
}

func (rlc *RateLimitedCalculator) Power(a, b int) (int, error) {
	allowed, err := rlc.rateLimiter.AllowRequest()

	if err != nil {
		return 0, err
	}

	if !allowed {
		return 0, err
	}

	result := 1
	for i := 0; i < b; i++ {
		result *= a
	}

	return result, nil
}

func (rlc *RateLimitedCalculator) Factorial(a int) (int, error) {
	allowed, err := rlc.rateLimiter.AllowRequest()

	if err != nil {
		return 0, err
	}

	if !allowed {
		return 0, err
	}

	result := 1
	for i := 1; i <= a; i++ {
		result *= i
	}

	return result, nil
}
