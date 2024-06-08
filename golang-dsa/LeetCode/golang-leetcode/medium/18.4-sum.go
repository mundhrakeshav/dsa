package leetcode_medium

import "sort"

/*
 * @lc app=leetcode id=18 lang=golang
 *
 * [18] 4Sum
 */

// @lc code=start

func FourSum(nums []int, target int) [][]int {
	return fourSum(nums, target)
}

func fourSum(nums []int, target int) [][]int {
	output := make([][]int, 0)
	n := len(nums)
	sort.Ints(nums)
	for i := 0; i < n; i++ {

		if i != 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			k := j + 1
			l := n - 1

			for k < l {
				sum := nums[i] + nums[j] + nums[k] + nums[l]
				if sum == target {
					output = append(output, []int{nums[i], nums[j], nums[k], nums[l]})
					k++
					l--
					for k < l && nums[k] == nums[k-1] {
						k++
					}
					for k < l && nums[l] == nums[l+1] {
						l--
					}
				} else if sum < target {
					k++
				} else {
					l--
				}

			}
		}
	}
	return output
}

// @lc code=end
