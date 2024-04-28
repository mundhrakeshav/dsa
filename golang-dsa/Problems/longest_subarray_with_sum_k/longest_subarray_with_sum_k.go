package longestsubarraywithsumk

// LongestSubArrayWithSumK: returns the length of the longest sub-array with sum k

func LongestSubArrayWithSumKHashMap(array []int64, k int64) struct {
	UpIdx, LowIdx int64
} {
	// Only Positives
	sumToIndexMap := make(map[int64]int64)
	var sum int64 = 0
	upIdx, lowIdx := int64(0), int64(0)
	for i, v := range array {
		idx := int64(i)
		sum += v

		if sum == k {
			lowIdx = 0
			upIdx = idx
		}

		rem := sum - k

		if prev_index, found := sumToIndexMap[rem]; found {
			if idx-prev_index > upIdx-lowIdx {
				upIdx = idx
				lowIdx = prev_index + 1
			}
		}

		if _, found := sumToIndexMap[sum]; !found {
			sumToIndexMap[sum] = idx
		}
	}
	return struct {
		UpIdx  int64
		LowIdx int64
	}{
		UpIdx:  upIdx,
		LowIdx: lowIdx,
	}
}

func LongestSubArrayWithSumK2Pointer(array []int64, k int64) struct {
	UpIdx, LowIdx int64
} {
	// Only Positives
	left, right := int64(0), int64(0)
	sum := array[0]
	upIdx, downIdx := int64(0), int64(0)
	n := int64(len(array))
	for right < n {
		for sum > k && left <= right {
			sum -= array[left]
			left++
		}

		if sum == k && upIdx-downIdx < right-left {
			downIdx = left
			upIdx = right
		}
		right++
		if right < n {
			sum += array[right]
		}
	}

	return struct {
		UpIdx  int64
		LowIdx int64
	}{
		UpIdx:  upIdx,
		LowIdx: downIdx,
	}
}
