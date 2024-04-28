package longestsubarraywithsumk

// LongestSubArrayWithSumK: returns the length of the longest sub-array with sum k

// [1,	1,	1,	2,	3,	2,	1, 	1, 	1, 	1, 1]

func LongestSubArrayWithSumKHashMap(array []int64, k int64) struct{
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
			if(idx - prev_index > upIdx - lowIdx) {
				upIdx = idx
				lowIdx = prev_index
			}
		}

		if _, found := sumToIndexMap[sum]; !found {
			sumToIndexMap[sum] = idx
		}
	}
	return struct{UpIdx int64; LowIdx int64}{
		UpIdx: upIdx,
		LowIdx: lowIdx,
	}
}
