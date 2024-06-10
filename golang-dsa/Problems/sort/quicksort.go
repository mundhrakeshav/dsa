package sort

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr) - 1)
}

func partition(arr []int, low, high int) int {
	pivot := arr[low]
	i := low
	j := high

	for i < j {
		for arr[i] <= pivot && i <= high-1 {
			i++
		}

		for arr[j] > pivot && j >= low+1 {
			j--
		}

		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[low], arr[j] = arr[j], arr[low] 
	return j
}

func quickSort(arr []int, low, high int) {
	if low < high {
		partitionIndex := partition(arr, low, high)
		quickSort(arr, 0, partitionIndex)
		quickSort(arr, partitionIndex+1, high)
	}
}
