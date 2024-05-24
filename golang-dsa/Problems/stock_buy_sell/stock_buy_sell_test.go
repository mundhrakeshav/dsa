package stockbuysell_test

import (
	"fmt"
	"testing"

	stockbuysell "github.com/mundhrakeshav/DSA/golang-dsa/Problems/stock_buy_sell"
)

func TestStockBuySellMultipleCases(t *testing.T) {
	tests := []struct {
		input    []int64
		expected int64
	}{
		{[]int64{7, 1, 5, 3, 6, 4}, 5},
		{[]int64{7, 6, 4, 3, 1}, 0},
		{[]int64{2, 4, 1}, 2},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Input: %v", test.input), func(t *testing.T) {
			actual, _ := stockbuysell.StockBuySell(test.input)
			if actual != test.expected {
				t.Errorf("Expected: %d, Got: %d", test.expected, actual)
			}
		})
	}
}
