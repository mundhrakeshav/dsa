package twosum_test

import (
	"fmt"
	"testing"

	twosum "github.com/mundhrakeshav/DSA/golang-dsa/Problems/2_sum"
)

func Test2Sum(t *testing.T) {
	arr := []int{2, 6, 5, 8, 11}
	fmt.Println(twosum.TwoSum(
		arr, 14,
	))
}
