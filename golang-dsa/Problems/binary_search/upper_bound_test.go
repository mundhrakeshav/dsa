package binarysearch_test

import (
	"fmt"
	"testing"

	binarysearch "github.com/mundhrakeshav/DSA/golang-dsa/Problems/binary_search"
)

func TestUpperBound(t *testing.T)  {
	// fmt.Println(binarysearch.UpperBound([]int{1, 3, 5, 6, 7, 8, 9, 11, 12, 15, 21, 24, 25, 27, 29, 31}, 25))
	fmt.Println(binarysearch.UpperBound([]int{1, 2, 2, 4}, 4))
}