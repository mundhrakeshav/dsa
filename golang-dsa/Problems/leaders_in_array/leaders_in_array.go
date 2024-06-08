package leadersinarray

import "math"

func LeadersInArray(arr []int) []int {
	n := len(arr)
	ret := make([]int, 0)
	max := math.MinInt
	for i := n - 1; i >= 0; i-- {
		if max < arr[i] {
			max = arr[i]
			ret = append(ret, arr[i])
		}
	}
	return ret
}
