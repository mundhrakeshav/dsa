package easy

import "sort"

/*
 * @lc app=leetcode id=1051 lang=golang
 *
 * [1051] Height Checker
 */

// @lc code=start
func heightChecker(heights []int) int {
    temp := make([]int, 0)
	temp = append(temp, heights...)

	sort.Ints(temp)
	count := 0
	for i, v := range heights {
		if temp[i] != v {
			count++
		}
	}
	return count
}
// @lc code=end

