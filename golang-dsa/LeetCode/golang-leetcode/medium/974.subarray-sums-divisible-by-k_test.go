package medium_test

import (
	"fmt"
	"testing"

	medium "github.com/mundhrakeshav/DSA/golang-dsa/LeetCode/golang-leetcode/medium"
)

func TestSubarraysDivByK(t *testing.T) {
	tests := []struct {
		input  []int
		target int
	}{
		{
			input:  []int{4, 5, 0, -2, -3, 1},
			target: 5,
		},
		{
			input:  []int{5, 0, 0, 0},
			target: 3,
		},
		{
			input:  []int{23, 2, 4, 6, 6},
			target: 7,
		},
		{
			input:  []int{23, 2, 4, 6, 7},
			target: 6,
		},
	}

	for _, test := range tests {
		t.Run("Test", func(t *testing.T) {
			fmt.Println(medium.SubarraysDivByK(test.input, test.target))
		})
	}

}
