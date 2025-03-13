package medium

/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 */

 func Search(nums []int, target int) int {
	 return search(nums, target)
 }

// @lc code=start
func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) / 2)
		switch {
		case nums[mid] == target:
			{
				return mid
			}

		case nums[low] <= nums[mid]:
			{
				// nums[low]..target..nums[mid]
				if nums[low] <= target && target <= nums[mid] {
					// target on left
					high = mid - 1
				} else {
					// target on right
					low = mid + 1
				}
			}
		default:
			{
				if nums[mid] <= target && target <= nums[high]{
					// target on right
					low = mid + 1
					} else {
						// target on left
					high = mid - 1
				}
				
			}

		}
	}
	return -1
}
// @lc code=end

