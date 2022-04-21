package selection_sort_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func findSmallest(items []int) int {
	smallest_index := 0
	smallest := items[0]

	for i := 1; i < len(items); i++ {
		current := items[i]
		if current < smallest {
			smallest_index = i
			smallest = current
		}
	}

	return smallest_index
}

func SelectionSort(items []int) []int {
	sorted := make([]int, len(items))

	for i, _ := range items {
		smallest := findSmallest(items)
		sorted[i] = items[smallest]

		items = append(items[:smallest], items[smallest+1:]...)
	}

	return sorted
}

func TestSelectionSort(t *testing.T) {
	assert.Equal(t, SelectionSort([]int{5, 3, 6, 2, 10}), []int{2, 3, 5, 6, 10}, "be ordered")
	assert.Equal(t, SelectionSort([]int{1, 2, 3, 4, 5}), []int{1, 2, 3, 4, 5}, "be ordered")
	assert.Equal(t, SelectionSort([]int{5, 4, 3, 2, 1}), []int{1, 2, 3, 4, 5}, "be ordered")
}
