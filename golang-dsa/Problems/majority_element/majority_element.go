package majorityelement

func MajorityElement2(arr []int) int {
	count, element := 0, 0
	for _, val := range arr {
		if count == 0 {
			element = val
		}
		if val == element {
			count++
		} else {
			count--
		}
	}
	count = 0
	for _, v := range arr {
		if v == element {
			count++
		}
	}
	if count > len(arr)/2 {
		return element
	}
	return -1
}

func MajorityElement3(arr []int) [2]int {
	count1, count2, element1, element2 := 0, 0, 0, 0
	for _, v := range arr {
		switch {
		case count1 == 0 && v != element2:
			{
				count1++
				element1 = v
			}
		case count2 == 0 && v != element1:
			{
				count2++
				element2 = v
			}
		case v == element1:
			{
				count1++
			}
		case v == element2:
			{
				count2++
			}
		default:
			{
				count1--
				count2--
			}
		}
	}
	count1, count2 = 0, 0

	for _, v := range arr {
		if v == element1 {
			count1++
		} else if v == element2 {
			count2++
		}
	}

	ret := [2]int{-1, -1}
	if count1 > len(arr) / 3 {
		ret[0] = element1
	}
	
	if count2 > len(arr) / 3 {
		ret[1] = element2
	}

	return ret
}
