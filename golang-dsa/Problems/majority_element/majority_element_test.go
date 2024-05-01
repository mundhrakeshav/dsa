package majorityelement_test

import (
	"fmt"
	"testing"

	majorityelement "github.com/mundhrakeshav/DSA/golang-dsa/Problems/majority_element"
)

func TestMajorityElement(t *testing.T) {
	arr := []int{2, 2, 1, 1, 1, 1, 1, 2, 2}
	fmt.Println(majorityelement.MajorityElement(arr))
}
