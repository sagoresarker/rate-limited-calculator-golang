package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockRateLimiter struct {
	allow bool
	err   error
}

func (m *MockRateLimiter) AllowRequest() (bool, error) {
	return m.allow, m.err
}

func TestRateLimitedCalculator(t *testing.T) {
	mockLimiter := &MockRateLimiter{allow: true, err: nil}
	calculator := NewRateLimitedCalculator(mockLimiter)

	t.Run("Test Add", func(t *testing.T) {
		result, err := calculator.Add(1, 2)
		assert.NoError(t, err)
		assert.Equal(t, 3, result)
	})

	t.Run("Test Subtract", func(t *testing.T) {
		result, err := calculator.Subtract(5, 2)
		assert.NoError(t, err)
		assert.Equal(t, 3, result)
	})

	t.Run("Test Multiply", func(t *testing.T) {
		result, err := calculator.Multiply(2, 3)
		assert.NoError(t, err)
		assert.Equal(t, 6, result)
	})

	t.Run("Test Divide", func(t *testing.T) {
		result, err := calculator.Divide(6, 2)
		assert.NoError(t, err)
		assert.Equal(t, 3, result)
	})

	t.Run("Test Modulo", func(t *testing.T) {
		result, err := calculator.Modulo(5, 2)
		assert.NoError(t, err)
		assert.Equal(t, 1, result)
	})

	t.Run("Test Power", func(t *testing.T) {
		result, err := calculator.Power(2, 3)
		assert.NoError(t, err)
		assert.Equal(t, 8, result)
	})

	t.Run("Test Factorial", func(t *testing.T) {
		result, err := calculator.Factorial(5)
		assert.NoError(t, err)
		assert.Equal(t, 120, result)
	})
}
