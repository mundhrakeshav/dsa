package longestconsecutivesubsequence_test

import (
	"fmt"
	"testing"

	longestconsecutivesubsequence "github.com/mundhrakeshav/DSA/golang-dsa/Problems/longest_consecutive_subsequence"
)

func TestLongestConsecutiveSubSeq(t *testing.T) {
	fmt.Println(longestconsecutivesubsequence.LongestConsecutiveSub([]int{100, 200, 1, 2, 3, 4}))
}
