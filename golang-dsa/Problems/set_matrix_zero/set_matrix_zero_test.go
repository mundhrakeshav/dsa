package setmatrixzero_test

import (
	"fmt"
	"testing"

	setmatrixzero "github.com/mundhrakeshav/DSA/golang-dsa/Problems/set_matrix_zero"
)

func TestSetMatrixZero(t *testing.T) {
	a := [][]int{
		{1, 1, 1}, {1, 0, 1}, {1, 1, 1},
	}
	fmt.Println(setmatrixzero.SetMatrixZero(a))
}
