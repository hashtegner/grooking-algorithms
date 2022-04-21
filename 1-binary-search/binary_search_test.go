package binary_search_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func BinarySearch(items []int, n int) int {
	low := 0
	high := len(items) - 1
	for low <= high {
		mid := (low + high) / 2

		item := items[mid]
		if item == n {
			return mid
		}

		if item > n {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

func TestBinarySearch(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}

	assert.Equal(t, BinarySearch(items, 1), 0, "1 to be present")
	assert.Equal(t, BinarySearch(items, 3), 2, "3 to be present")
	assert.Equal(t, BinarySearch(items, 5), 4, "5 to be present")
	assert.Equal(t, BinarySearch(items, 6), -1, "6 to not be present")
	assert.Equal(t, BinarySearch(items, 0), -1, "0 to not be present")
}
