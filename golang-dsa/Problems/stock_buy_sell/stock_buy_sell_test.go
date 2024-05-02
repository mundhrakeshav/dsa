package stockbuysell_test

import (
	"fmt"
	"testing"

	stockbuysell "github.com/mundhrakeshav/DSA/golang-dsa/Problems/stock_buy_sell"
)

func TestStockBuySell(t *testing.T) {
	fmt.Println(stockbuysell.StockBuySell([]int64{
		7, 1, 5, 3, 6, 4,
	}))
}
