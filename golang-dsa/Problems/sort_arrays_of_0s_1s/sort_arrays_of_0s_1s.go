package sortarraysof0s1s

func SortArrayOf0s1s2s(arr []int) {
	n := len(arr)
	low, mid, high := 0, 0, n-1

	for mid <= high {
		switch {
		case arr[mid] == 0:
			{
				// temp := arr[mid]
				// arr[mid] = arr[low]
				// arr[low] = temp
				arr[mid], arr[low] = arr[low], arr[mid]
				low++
				mid++
			}
		case arr[mid] == 1:
			{
				mid++
			}
		case arr[mid] == 2:
			{
				// temp := arr[mid]
				// arr[mid] = arr[high]
				// arr[high] = temp
				arr[mid], arr[high] = arr[high], arr[mid]
				high--
			}
		}
	}

}
