package longestsubarraywithsumk_test

import (
	"fmt"
	"testing"

	longestsubarraywithsumk "github.com/mundhrakeshav/DSA/golang-dsa/Problems/longest_subarray_with_sum_k"
)

func TestLongestSubArrayWithSumK(t *testing.T) {
	arr := []int64{
		1, 1, 1, 2, 3, 2, 1, 1, 1, 1, 1,
	}
	indexes := longestsubarraywithsumk.LongestSubArrayWithSumKHashMap(arr, 8)
	fmt.Print(indexes.UpIdx, indexes.LowIdx)
}
