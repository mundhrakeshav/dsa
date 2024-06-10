package medium_test

import (
	"fmt"
	"testing"

	medium "github.com/mundhrakeshav/DSA/golang-dsa/LeetCode/golang-leetcode/medium"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		input []int
	}{
		{
			input: []int{-1, 0, 1, 2, -1, -4},
		},
		{
			input: []int{0, 0, 0, 0},
		},
	}

	for _, test := range tests {
		t.Run("Test", func(t *testing.T) {
			fmt.Println(medium.ThreeSum(test.input))
		})
	}
}
