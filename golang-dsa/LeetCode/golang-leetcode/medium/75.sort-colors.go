package medium

/*
 * @lc app=leetcode id=75 lang=golang
 *
 * [75] Sort Colors
 */

// @lc code=start
func sortColors(nums []int) {
	left := 0
	right := len(nums) - 1

	for i := 0; i < len(nums) && i <= right; i++ {
		if nums[i] == 0 {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		} else if nums[i] == 2 {
			nums[i], nums[right] = nums[right], nums[i]
			right--
			i--
		}
	}

}

// 2Pass Approach
// func sortColors(nums []int) {
// 	zeros, ones, twos := 0, 0, 0
// 	for _, v := range nums {
// 		switch v {
// 		case 0:
// 			zeros++
// 		case 1:
// 			ones++
// 		default:
// 			twos++
// 		}
// 	}

// 	for i := 0; i < len(nums); i++ {
// 		switch {
// 		case i < zeros:
// 			nums[i] = 0
// 		case i >= zeros && i < ones+zeros:
// 			nums[i] = 1
// 		default:
// 			nums[i] = 2
// 		}
// 	}
// }

// @lc code=end
