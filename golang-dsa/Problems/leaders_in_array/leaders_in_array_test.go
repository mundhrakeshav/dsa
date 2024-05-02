package leadersinarray_test

import (
	"fmt"
	"testing"

	leadersinarray "github.com/mundhrakeshav/DSA/golang-dsa/Problems/leaders_in_array"
)

func TestLeadersInArray(t *testing.T) {
	fmt.Println(leadersinarray.LeadersInArray([]int{10, 22, 12, 3, 0, 6}))
}
