package binarysearch

import (
	"testing"
)


func TestFloor(t *testing.T) {
	tests := []struct {
		arr      []int
		num      int
		expected int
	}{
		{[]int{3, 4, 4, 7, 8, 10}, 8, 8},
		{[]int{3, 4, 4, 7, 8, 10}, 9, 8},
		{[]int{1, 2, 4, 6, 8, 10}, 5, 4},
		{[]int{1, 2, 4, 6, 8, 10}, 10, 10},
		{[]int{1, 2, 4, 6, 8, 10}, 1, 1},
		{[]int{1, 2, 4, 6, 8, 10}, 11, 10},
		{[]int{1, 2, 4, 6, 8, 10}, 0, -1},
	}

	for _, test := range tests {
		result := floor(test.arr, test.num)
		if result != test.expected {
			t.Errorf("For array %v and number %d, expected floor index %d, but got %d", test.arr, test.num, test.expected, result)
		}
	}
}

func TestCeil(t *testing.T) {
	tests := []struct {
		arr      []int
		num      int
		expected int
	}{
		{[]int{3, 4, 4, 7, 8, 10}, 8, 8},
		{[]int{3, 4, 4, 7, 8, 10}, 9, 10},
		{[]int{1, 2, 4, 6, 8, 10}, 5, 6},
		{[]int{1, 2, 4, 6, 8, 10}, 10, 10},
		{[]int{1, 2, 4, 6, 8, 10}, 1, 1},
		{[]int{1, 2, 4, 6, 8, 10}, 11, -1},
		{[]int{1, 2, 4, 6, 8, 10}, 0, 1},
	}

	for _, test := range tests {
		result := ceil(test.arr, test.num)
		if result != test.expected {
			t.Errorf("For array %v and number %d, expected ceil value %d, but got %d", test.arr, test.num, test.expected, result)
		}
	}
}
