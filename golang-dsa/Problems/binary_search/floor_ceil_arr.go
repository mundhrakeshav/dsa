package binarysearch

// // Problem Statement: You're given an sorted array arr of n integers and an integer x. Find the floor and ceiling of x in arr[0..n-1].
// // The floor of x is the largest element in the array which is smaller than or equal to x.
// // The ceiling of x is the smallest element in the array greater than or equal to x.
func FloorAndCeilOfArray(arr []int, num int) (f, c int) {
	f = floor(arr, num)
	c = ceil(arr, num)

	return f, c
}


func floor(arr []int, num int) int {
	low, high := 0, len(arr)-1
	floor := -1 
	for low <= high {
		mid := low + (high-low)/2

		switch {
		case arr[mid] <= num:
			{
				floor = arr[mid]  // Store the value instead of the index
				low = mid + 1
			}
		default:
			{
				high = mid - 1
			}
		}
	}

	return floor
}

func ceil(arr []int, num int) int {
	low, high := 0, len(arr)-1
	ceil := -1

	for low <= high {
		mid := low + (high - low) / 2

		switch {
		case arr[mid] < num :
			low = mid + 1
		default : 
			high = mid - 1
			ceil = arr[mid]
		}
	}

	return ceil
}

