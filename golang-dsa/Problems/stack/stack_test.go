package stack_test

import (
	"fmt"
	"testing"

	"github.com/mundhrakeshav/DSA/golang-dsa/Problems/stack"
)


func TestNGE(t *testing.T) {
	arr := []int{3,10,4,2,1,2,6,1,7,2,9}
	fmt.Println(stack.NGEs(arr))
}
