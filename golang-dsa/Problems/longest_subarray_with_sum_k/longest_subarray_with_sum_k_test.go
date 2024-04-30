package longestsubarraywithsumk_test

import (
	"fmt"
	"testing"

	longestsubarraywithsumk "github.com/mundhrakeshav/DSA/golang-dsa/Problems/longest_subarray_with_sum_k"
)

// [
// 	1,
// 	1,
// 	1,
// 	2,
// 	3,
// 	2,
// 	1,
// 	1,
// 	1,
// 	1,
// 	1
// ]

func TestLongestSubArrayWithSumK(t *testing.T) {
	// arr := []int64{
	// 	1, 1, 1, 2, 3, 2, 1, 1, 1, 1, 1,
	// }
	arr := []int64{
		-1, 1, 1, -3, 11, 3, -11,
	}
	indexes := longestsubarraywithsumk.LongestSubArrayWithSumKHashMap(arr, 1)
	fmt.Print(indexes.UpIdx, indexes.LowIdx)
}

func TestLongestSubArrayWithSumK2Pointer(t *testing.T) {
	arr := []int64{
		-1, 1, 1, -3, 11, 3, -11,
	}
	indexes := longestsubarraywithsumk.LongestSubArrayWithSumK2Pointer(arr, 1)
	fmt.Print(indexes.UpIdx, indexes.LowIdx)
}
