package majorityelement

func MajorityElement(arr []int) (int){
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
	if count > len(arr) / 2 {
		return element
	}
	return -1
}