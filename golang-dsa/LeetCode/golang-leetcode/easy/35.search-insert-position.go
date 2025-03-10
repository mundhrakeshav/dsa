package easy

/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1
	idx := high

	if target < nums[low] {
		return 0
	}

	if target > nums[high] {
		return high + 1
	}

	for low <= high {
		mid := low + (high-low)/2
		switch {
		case nums[mid] >= target:
			idx = mid
			high = mid - 1
		default:
			low = mid + 1
		}
	}

	return idx
}

// @lc code=end
