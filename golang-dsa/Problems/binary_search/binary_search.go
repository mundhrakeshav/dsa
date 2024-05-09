package binarysearch

func BinarySearchIterative(arr []int, target int) int {
	return binarySearchIterative(arr, 0, len(arr)-1, target)
}

func BinarySearchRecursive(arr []int, target int) int {
	return binarySearchRecursive(arr, 0, len(arr)-1, target)
}

func binarySearchIterative(arr []int, low, high, target int) int {
	for low <= high {
		mid := low + (high - low/2)
		switch {
		case arr[mid] == target:
			{
				return mid
			}
		case arr[mid] > target:
			{
				high = mid - 1
			}
		default:
			{
				low = mid + 1
			}
		}
	}
	return -1
}

func binarySearchRecursive(arr []int, low, high, target int) int {
	if low > high {
		return -1
	}

	mid := low + (high - low/2)
	switch {
	case arr[mid] == target:
		{
			return mid
		}
	case arr[mid] > target:
		{
			return binarySearchRecursive(arr, low, mid-1, target)
		}
	default:
		{
			return binarySearchRecursive(arr, mid+1, high, target)
		}
	}

}
