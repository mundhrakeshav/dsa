package maxsubarray_test

import (
	"fmt"
	"testing"

	maxsubarray "github.com/mundhrakeshav/DSA/golang-dsa/Problems/max_subarray"
)

func TestMaxSubArray(t *testing.T) {
	fmt.Println(maxsubarray.MaxSubArray([]int{-2,-3,4,-1,-2,1,5,-3}))
}
