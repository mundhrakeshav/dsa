package medium

/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	f := first(nums, target)
	l := last(nums, target)
	return []int{f, l}
}

func first(nums []int, target int) int {
	low, high := 0, len(nums) - 1
	first := -1
	for low <= high {
		mid := low + ((high - low) / 2)
		switch {
		case nums[mid] == target: {
			first= mid
			high = mid - 1
		}
		case nums[mid] < target: {
			low = mid + 1
		}
		case nums[mid] > target: {
			high = mid - 1
		}
		
		
		}
	}

	return first
}
func last(nums []int, target int) int {
	low, high := 0, len(nums) - 1
	last := -1
	for low <= high {
		mid := low + ((high - low) / 2)
		switch {
		case nums[mid] == target: {
			last= mid
			low = mid + 1
		}
		case nums[mid] < target: {
			low = mid + 1
		}
		case nums[mid] > target: {
			high = mid - 1
		}
		
		
		}
	}

	return last
}
// @lc code=end
