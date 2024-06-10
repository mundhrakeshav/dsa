package sort

func MergeSort(nums []int) {
	mergeSort(nums, 0, len(nums)-1)
}

func merge(nums []int, low, mid, high int) {
	tempArr := make([]int, 0)
	idx_1, idx_2 := low, mid+1

	for idx_1 <= mid && idx_2 <= high {
		if nums[idx_1] <= nums[idx_2] {
			tempArr = append(tempArr, nums[idx_1])
			idx_1++
		} else {
			tempArr = append(tempArr, nums[idx_2])
			idx_2++
		}
	}

	for idx_1 <= mid {
		tempArr = append(tempArr, nums[idx_1])
		idx_1++
	}

	for idx_2 <= high {
		tempArr = append(tempArr, nums[idx_1])
		idx_2++
	}

	for i := low; i <= high; i++ {
		nums[i] = tempArr[i-low]
	}

}

func mergeSort(nums []int, low, high int) {
	if low >= high {
		return
	}

	mid := (high + low) / 2

	mergeSort(nums, low, mid)
	mergeSort(nums, mid+1, high)
	merge(nums, low, mid, high)
}
