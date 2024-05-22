package ratelimiter

import (
	"fmt"
	"time"
)

var ErrRateLimitExceeded = fmt.Errorf("rate limit exceeded")

type RateLimiter interface {
	AllowRequest() (bool, error)
}
type MockRateLimiter struct {
	Allow bool
	Err   error
}

func (m *MockRateLimiter) AllowRequest() (bool, error) {
	return m.Allow, m.Err
}

type Storage interface {
	IncrementCount(windowStart time.Time, count int) error
}

type SlidingWindowRateLimiter struct {
	limit        int
	windowSize   time.Duration
	storage      Storage
	windowStart  time.Time
	currentCount int
}

func NewRateLimiter(limit int, storage Storage) *SlidingWindowRateLimiter {
	return &SlidingWindowRateLimiter{
		limit:        limit,
		storage:      storage,
		windowSize:   time.Minute,
		windowStart:  time.Now(),
		currentCount: 0,
	}
}

func (rl *SlidingWindowRateLimiter) AllowRequest() (bool, error) {
	now := time.Now()

	if now.Sub(rl.windowStart) >= rl.windowSize {
		rl.windowStart = now
		rl.currentCount = 0
	}

	if rl.currentCount >= rl.limit {
		return false, ErrRateLimitExceeded
	}

	rl.currentCount++

	err := rl.storage.IncrementCount(rl.windowStart, rl.currentCount)
	if err != nil {
		return false, err
	}

	return true, nil
}
