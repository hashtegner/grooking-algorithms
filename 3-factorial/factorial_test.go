package factorial_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * Factorial(n-1)
}

func TestFactorial(t *testing.T) {
	assert.Equal(t, Factorial(0), 1, "0 factorial should be 1")
	assert.Equal(t, Factorial(1), 1, "1 factorial should be 1")
	assert.Equal(t, Factorial(5), 120, "5 factorial should be 120")
}
