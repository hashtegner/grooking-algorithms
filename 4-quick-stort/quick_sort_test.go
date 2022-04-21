package quick_sort_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func QuickSort(items []int) []int {
	if len(items) < 2 {
		return items
	}

	pivot := items[0]
	smallest := make([]int, 0)
	greatest := make([]int, 0)

	for _, current := range items[1:] {
		if current > pivot {
			greatest = append(greatest, current)
		} else {
			smallest = append(smallest, current)
		}
	}

	sorted := append(QuickSort(smallest), pivot)
	sorted = append(sorted, QuickSort(greatest)...)

	return sorted
}

func TestQuickSort(t *testing.T) {
	assert.Equal(t, QuickSort([]int{10, 5, 2, 3}), []int{2, 3, 5, 10})

	// bad performance
	assert.Equal(t, QuickSort([]int{1, 2, 3, 4, 5}), []int{1, 2, 3, 4, 5})
	assert.Equal(t, QuickSort([]int{5, 4, 3, 2, 1}), []int{1, 2, 3, 4, 5})
}
