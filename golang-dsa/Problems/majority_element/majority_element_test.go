package majorityelement_test

import (
	"fmt"
	"testing"

	majorityelement "github.com/mundhrakeshav/DSA/golang-dsa/Problems/majority_element"
)

func TestMajorityElement2(t *testing.T) {
	arr := []int{2, 2, 1, 1, 1, 1, 1, 2, 2}
	fmt.Println(majorityelement.MajorityElement2(arr))
}
func TestMajorityElement3(t *testing.T) {
	arr := []int{2, 2, 1, 1, 1, 1, 1, 2, 3, 3, 2}
	fmt.Println(majorityelement.MajorityElement3(arr))
}
