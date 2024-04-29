package twosum

import (
	"slices"
)

func TwoSum(nums []int, target int) []int {
	
	n := len(nums)
	idx1, idx2 := int(0), n -1

	// {2,6,5,8,11}
	// Sort and 2 Pointer
	slices.Sort(nums)
	
	for idx2 < n {
		for idx1 < idx2 {
			switch {
				case nums[idx1] + nums[idx2] < target: {
					idx1++
				}
				case nums[idx1] + nums[idx2] > target: {
					idx2--
				}
				
				case nums[idx1] + nums[idx2] == target: {
					return []int{idx1, idx2}
				}

			}		
		}
	}
	return []int{-1, -1}
}
