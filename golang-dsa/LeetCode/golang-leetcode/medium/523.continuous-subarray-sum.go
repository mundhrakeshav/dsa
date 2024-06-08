package leetcode_medium

import "slices"

/*
 * @lc app=leetcode id=523 lang=golang
 *
 * [523] Continuous Subarray Sum
 */

// @lc code=start
func CheckSubarraySum(nums []int, k int) bool {
	return checkSubarraySum(nums, k)
}
func checkSubarraySum(nums []int, k int) bool {
	prefixMod := make(map[int]int)
	prefixMod[0] = -1
	mod := 0
	for i, v := range nums {
		mod = (mod + v) % k

		if prevIndex, found := prefixMod[mod]; found  {
			if (i-prevIndex) > 1 {
				return true
			}
		} else {
			prefixMod[mod] = i
		}
	}
	return false
}

// @lc code=end


// ???????
func checkSubarraySum1(nums []int, k int) bool {
    if len(nums) <= 1{
        return false
    }
    isMonotonicallyDecreasing := true
    sumarr := make([]int,len(nums)+1)
    startSumIdx := 0
    for idx:=0;idx<len(nums);idx++ {
        sumarr[idx+1] = sumarr[idx] + nums[idx]
        if idx >= 1 && sumarr[idx+1] == sumarr[idx] && sumarr[idx+1] == sumarr[idx-1]{
            return true
        }
        if idx >= 1 && sumarr[idx+1]%k == 0 {
            return true
        }
        if sumarr[idx+1] < k {
            startSumIdx = idx+1
        }
        if idx >=1 && nums[idx] > nums[idx-1] {
            isMonotonicallyDecreasing = false
        }
    }
    if isMonotonicallyDecreasing {
        slices.Reverse(nums)
        for idx:=0;idx<len(nums);idx++ {
            sumarr[idx+1] = sumarr[idx] + nums[idx]
            if idx >= 1 && sumarr[idx+1] == sumarr[idx] && sumarr[idx+1] == sumarr[idx-1]{
                return true
            }
            if idx >= 1 && sumarr[idx+1]%k == 0 {
                return true
            }
            if sumarr[idx+1] < k {
                startSumIdx = idx+1
            }
            if idx >=1 && nums[idx] > nums[idx-1] {
                isMonotonicallyDecreasing = false
            }
        }
    }
    if sumarr[len(sumarr)-1] < k {
        return false
    }
    // nums arry =[23,2,6,4,7]
    // sumarr = [0,23,25,31,35,42]
    for i:=1;i<len(sumarr)-2;i++{
        j := i+2
        if startSumIdx > j {
            j = startSumIdx
        }
        for ;j<len(sumarr);j++{
            if (sumarr[j]-sumarr[i]) % k == 0 {
                return true
            }
        }
    }
    return false


    //https://leetcode.com/problems/continuous-subarray-sum/solutions/99499/java-o-n-time-o-k-space/
    // (fastest)
    // prevsum := 0
    // sumpast := map[int]int{0:-1}
    // for i := range nums {
    //     new := (prevsum + nums[i])%k
    //     prevsum = new
    //     if val, ok := sumpast[new]; ok {
    //         if i-val > 1 {
    //             return true
    //         }
    //     } else {
    //         sumpast[new] = i
    //     }
    // }
    // return false
    // prevSum := 0
    // sumArr := []int{0}
    // for i := range nums {
    //     prevSum += nums[i]
    //     sumArr = append(sumArr, prevSum)
    // } 
    // for i := 0;i<len(sumArr)-1;i++{
    //     for j:= i+2;j<len(sumArr);j++ {
    //         if (sumArr[j]-sumArr[i])%k == 0 {
    //             return true
    //         }
    //     }
    // }
    // return false
}