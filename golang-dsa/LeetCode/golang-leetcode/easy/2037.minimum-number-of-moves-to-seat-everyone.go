package easy

import (
	"sort"
)

/*
 * @lc app=leetcode id=2037 lang=golang
 *
 * [2037] Minimum Number of Moves to Seat Everyone
 */

// @lc code=start
func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)

	sum := 0

	for i := 0; i < len(seats); i++ {
		sum += absDiffInt(seats[i], students[i])
	}
	return sum
}

func absDiffInt(x, y int) int {
	if x < y {
	   return y - x
	}
	return x - y
 }
// @lc code=end
