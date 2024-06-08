package stockbuysell

import "math"

func StockBuySell(arr []int64) (int64, [2]int64) {
	maxPro := int64(0)
	min := int64(math.MaxInt64)
	i1, i2 := int64(0), int64(0)
	for i, v := range arr {
		if v < min {
			i1 = int64(i)
			min = v
		}
		if v-min > maxPro {
			i2 = int64(i)
			maxPro = v - min
		}
	}
	return maxPro, [2]int64{i1, i2}
}
