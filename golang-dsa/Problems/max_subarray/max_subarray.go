package maxsubarray

func MaxSubArray(arr []int) (int, [2]int) {
	sum, max_sum := 0, 0
	// -2, 1, -3, 4, -1, 2, 1, -5, 4
	i1, i2 := -1, -1

	for i, val := range arr {
		if sum == 0 {
			i1 = i
		}
		sum += val
		if sum > max_sum {
			max_sum = sum
			i2 = i
		}
		if sum < 0 {
			sum = 0
		}
	}
	return max_sum, [2]int{i1, i2}
}
