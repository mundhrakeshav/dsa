package longestconsecutivesubsequence

func LongestConsecutiveSub(arr []int) int {
	count, max_count := 0, 0
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] == 1 {
			count++
			if count > max_count {
				max_count = count
			}
		} else {
			count = 1
		}
	}
	return max_count
}
