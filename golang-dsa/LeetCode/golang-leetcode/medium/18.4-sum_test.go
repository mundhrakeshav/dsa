package leetcode_medium_test

import (
	"fmt"
	"testing"

	leetcode_medium "github.com/mundhrakeshav/DSA/golang-dsa/LeetCode/golang-leetcode/medium"
)

func TestFourSum(t *testing.T) {
	tests := []struct {
		input  []int
		target int
	}{
		{
			input:  []int{1, 0, -1, 0, -2, 2},
			target: 0,
		},
		{
			input:  []int{-2, -1, -1, 1, 1, 2, 2},
			target: 0,
		},
		{
			input:  []int{2, 2, 2, 2, 2},
			target: 8,
		},
	}

	for _, test := range tests {
		t.Run("Test", func(t *testing.T) {
			fmt.Println(leetcode_medium.FourSum(test.input, test.target))
		})
	}
}
