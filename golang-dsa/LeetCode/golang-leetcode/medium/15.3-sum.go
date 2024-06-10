package medium

import (
	"sort"
)

/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start
func ThreeSum(nums []int) [][]int {
	return threeSum(nums)
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	output := make([][]int, 0)
	last_element := len(nums) - 1
	for i, v := range nums {

		// No need to check numbers after 0. As they are already sorted and
		// sum can't be zero
		// v > 0 -> break
		if v > 0 {
			break
		}
		if i != 0 && v == nums[i-1] {
			continue
		}

		j, k := i+1, last_element

		for j < k {
			sum := v + nums[j] + nums[k]

			if sum == 0 {
				output = append(output, []int{v, nums[j], nums[k]})
				j++
				k--

				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}

	}
	return output
}

// @lc code=end
