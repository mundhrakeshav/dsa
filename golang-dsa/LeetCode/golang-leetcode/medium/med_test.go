package medium_test

import (
	"testing"

	"github.com/mundhrakeshav/DSA/golang-dsa/LeetCode/golang-leetcode/medium"
)

func TestRotatedSortedSearchUnique(t *testing.T) {
    tests := []struct {
        nums     []int
        target   int
        expected int
    }{
        // Standard cases
        {[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},      // Target in second part of rotated array
        {[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},     // Target not in array
        {[]int{1, 3, 5}, 5, 2},                  // Non-rotated array
        {[]int{4, 5, 6, 7, 8, 1, 2, 3}, 8, 4},   // Target just before rotation point
        {[]int{4, 5, 6, 7, 8, 1, 2, 3}, 1, 5},   // Target just after rotation point
        
        // Edge cases
        {[]int{}, 5, -1},                        // Empty array
        {[]int{5}, 5, 0},                        // Single element array with target
        {[]int{5}, 6, -1},                       // Single element array without target
        {[]int{1, 3}, 1, 0},                     // Two-element array, target is first
        {[]int{3, 1}, 1, 1},                     // Two-element array, target is second (rotated)
        {[]int{3, 1}, 3, 0},                     // Two-element array, target is first (rotated)
        
        // More rotation cases
        {[]int{7, 8, 1, 2, 3, 4, 5, 6}, 2, 3},   // Target in middle after rotation
        {[]int{6, 7, 8, 9, 1, 2, 3, 4, 5}, 6, 0}, // Target at first position
        {[]int{6, 7, 8, 9, 1, 2, 3, 4, 5}, 5, 8}, // Target at last position
        {[]int{1}, 0, -1},                        // Single element, target not found
        {[]int{3, 4, 5, 6, 7, 8, 9, 1, 2}, 3, 0}, // Target at start before rotation
        {[]int{5, 1, 2, 3, 4}, 1, 1},            // Target right after rotation point
    }

    for _, test := range tests {
        result := medium.Search(test.nums, test.target)
        if result != test.expected {
            t.Errorf("For array %v and target %d, expected index %d, but got %d", 
                test.nums, test.target, test.expected, result)
        }
    }
}