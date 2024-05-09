package binarysearch

import (
	"fmt"
)
func UpperBound(arr []int, x int) int {
	fmt.Println(len(arr))
	return upperBoundIter(arr, 0, len(arr)-1, x)
}
func upperBoundIter(arr []int, low, high, x int) int {
	lowestIndex := high
	for low <= high {
		mid := low + (high - low/2)
		switch {
		case arr[mid] > x:
			{
				if mid < lowestIndex {
					lowestIndex = mid
				}
				high = mid - 1
			}
		default:
			{
				low = mid + 1
			}
		}
	}
	return lowestIndex
}
