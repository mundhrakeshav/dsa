package sortarraysof0s1s_test

import (
	"fmt"
	"testing"

	sortarraysof0s1s "github.com/mundhrakeshav/DSA/golang-dsa/Problems/sort_arrays_of_0s_1s"
)

func TestSortArraysOf0s1s2s(t *testing.T) {
	arr := []int{2, 1, 0, 1, 1, 0, 1, 2, 1, 1, 2, 0, 0, 0}
	sortarraysof0s1s.SortArrayOf0s1s2s(arr)
	fmt.Println(arr)
}
