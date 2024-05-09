package binarysearch_test

import (
	"fmt"
	"testing"

	binarysearch "github.com/mundhrakeshav/DSA/golang-dsa/Problems/binary_search"
)

func TestBinarySearchIterative(t *testing.T) {
	fmt.Println(binarysearch.BinarySearchIterative([]int{1, 3, 5, 6, 7, 8, 9, 11, 12, 15, 21, 24, 25, 27, 29, 31}, 56))
}
func TestBinarySearchRecursive(t *testing.T) {
	fmt.Println(binarysearch.BinarySearchRecursive([]int{1, 3, 5, 6, 7, 8, 9, 11, 12, 15, 21, 24, 25, 27, 29, 31}, 45))
}
