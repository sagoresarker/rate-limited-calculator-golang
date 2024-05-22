package ratelimiter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type MockStorage struct {
	counts map[int64]int
}

func (m *MockStorage) IncrementCount(windowStart time.Time, count int) error {
	if m.counts == nil {
		m.counts = make(map[int64]int)
	}
	m.counts[windowStart.Unix()] = count
	return nil
}

func TestSlidingWindowRateLimiter(t *testing.T) {
	mockStorage := &MockStorage{}
	rateLimiter := NewRateLimiter(10, mockStorage)

	for i := 0; i < 10; i++ {
		allowed, err := rateLimiter.AllowRequest()
		assert.NoError(t, err)
		assert.True(t, allowed)
	}

	allowed, err := rateLimiter.AllowRequest()
	assert.Error(t, err)
	assert.False(t, allowed)
	assert.Equal(t, ErrRateLimitExceeded, err)
}
