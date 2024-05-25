package graph_test

import (
	"fmt"
	"testing"

	"github.com/mundhrakeshav/DSA/golang-dsa/Problems/graph"
)

func TestRottenOranges(t *testing.T) {
	fmt.Println(graph.RottenOranges([][]int{
		{2, 1, 1},
		{0, 1, 1},
		{0, 0, 1},
	}))
}
