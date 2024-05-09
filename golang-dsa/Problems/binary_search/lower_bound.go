package binarysearch

import "fmt"

func LowerBound(arr []int, x int) int {
	// find the smallest index such that arr[index] >= x
	fmt.Println(len(arr))
	return lowerBoundIter(arr, 0, len(arr)-1, x)
}
func lowerBoundIter(arr []int, low, high, x int) int {
	lowestIndex := high
	for low <= high {
		mid := low + (high - low/2)
		switch {
		case arr[mid] >= x:
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
