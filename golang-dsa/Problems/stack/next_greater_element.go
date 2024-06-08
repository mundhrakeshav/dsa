package stack

import (
	"github.com/golanglibs/gocollections/stack/arraystack"
)

func NGEs(arr []int) []int {
	stk := arraystack.New[int]()
	n := len(arr)
	nge := make([]int, n)

	for i := 2*n - 1; i >= 0; i-- {

		for !stk.Empty() && *stk.Peek() <= arr[i%n] {
			stk.Pop()
		}

		if i < n {
			if stk.Empty() {
				nge[i] = -1
			} else {
				nge[i] = *stk.Peek()
			}
		}

		stk.Push(arr[i%n])
	}

	return nge
}
