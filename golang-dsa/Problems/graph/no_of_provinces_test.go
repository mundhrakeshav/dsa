package graph_test

import (
	"testing"

	"github.com/mundhrakeshav/DSA/golang-dsa/Problems/graph"
)

func TestGetProvincesCount(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected int
	}{
		// {"Test Case 1", [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}, 2},
		{"Test Case 2", [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := graph.GetProvincesCountRecursive(tt.input)
			if actual != tt.expected {
				t.Errorf("For input %v, expected %d, but got %d", tt.input, tt.expected, actual)
			}
		})
	}
}

func TestGetProvincesCountIterative(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected int
	}{
		{"Test Case 1", [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}, 2},
		{"Test Case 2", [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := graph.GetProvincesCountIterative(tt.input)
			if actual != tt.expected {
				t.Errorf("For input %v, expected %d, but got %d", tt.input, tt.expected, actual)
			}
		})
	}
}
