package sort_test

import (
	"fmt"
	"testing"

	"github.com/mundhrakeshav/DSA/golang-dsa/Problems/sort"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		input  []int
		output int
	}{
		{
			input: []int{5, 4, 0, -2, -3, 1},
		},

		{
			input: []int{23, 2, 4, 6, 7},
		},
	}
	for _, test := range tests {
		t.Run("Test", func(t *testing.T) {
			sort.QuickSort(test.input)
			fmt.Println(test.input)
		})
	}

}
