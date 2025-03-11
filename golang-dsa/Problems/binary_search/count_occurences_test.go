package binarysearch_test

import (
	"testing"

	binarysearch "github.com/mundhrakeshav/DSA/golang-dsa/Problems/binary_search"
)

func TestCountOccurences(t *testing.T) {
	tests := []struct {
		arr      []int
		target   int
		expected int
	}{
		{[]int{2, 2, 3, 3, 3, 3, 4}, 3, 4},
		{[]int{1, 1, 2, 2, 2, 2, 2, 3}, 2, 5},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5, 1},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 10, 0},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1}, 1, 9},
		{[]int{}, 1, 0},

		// Edgy
		{[]int{3}, 3, 1},                         // Single element array with target
		{[]int{5}, 6, 0},                         // Single element array without target
		{[]int{2, 2, 2, 2, 2}, 2, 5},             // All elements are target
		{[]int{-5, -5, -5, -4, -3}, -5, 3},       // Negative numbers
		{[]int{3, 3, 3, 4, 5, 6}, 3, 3},          // Target only at beginning
		{[]int{1, 2, 3, 5, 5, 5}, 5, 3},          // Target only at end
		{[]int{5, 6, 7, 8, 9, 10}, 4, 0},         // Target smaller than all elements
		{[]int{1, 2, 3, 4, 5, 6}, 7, 0},          // Target larger than all elements
		{[]int{5, 10, 10, 10, 10, 20}, 10, 4},    // Target in middle with repetition
		{[]int{1, 1, 1, 2, 3, 3, 3}, 3, 3},       // Target at end with repetition
		{[]int{3, 3, 3, 4, 5, 6, 7, 8, 9}, 3, 3}, // Multiple occurrences at start
		{[]int{1, 2, 3, 4, 5, 5, 5}, 5, 3},       // Multiple occurrences at end
		{[]int{1, 2, 3, 3, 3, 3, 4, 5}, 3, 4},    // Multiple occurrences in middle
		{[]int{2147483647}, 2147483647, 1},       // Max integer value
		{[]int{-2147483648}, -2147483648, 1},     // Min integer value

	}

	for _, test := range tests {
		result := binarysearch.CountOccurences(test.arr, test.target)
		if result != test.expected {
			t.Errorf("For array %v and target %d, expected count %d, but got %d", test.arr, test.target, test.expected, result)
		}
	}
}
