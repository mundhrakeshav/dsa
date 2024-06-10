package easy

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	// value to index
	record := make(map[int]int, 0)

	for i, v := range nums {
		diff := target - v
		if diffIndex, found := record[diff]; found && i != diffIndex {
			return []int{i, diffIndex}
		}
		/*
			say corresponding element to v is x
			The x isn't found in map right now so we save current element's
			index. Further in iterations, x will be encountered so we as v_1
			and v + v_1 == target.
			Index of v then could be fetched from record
		*/
		record[v] = i
	}

	return []int{-1}
}

// @lc code=end
