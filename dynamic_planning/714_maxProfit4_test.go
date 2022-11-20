package dynamic_planning

import (
	"leetcode/utils"
	"testing"
)

func maxProfit4(prices []int, fee int) int {
	n := len(prices)
	dp0, dp1 := 0, -prices[0]
	for i := 1; i < n; i++ {
		dp0, dp1 = utils.MaxInt(dp0, dp1+prices[i]-fee), utils.MaxInt(dp1, dp0-prices[i])
	}
	return utils.MaxInt(dp0, dp1)
}

func TestMaxProfit4(t *testing.T) {
	tests := []struct {
		inputPrices []int
		inputFee    int
		want        int
	}{
		{[]int{1, 3, 2, 8, 4, 9}, 2, 8},
		{[]int{1, 3, 7, 5, 10, 3}, 3, 6},
	}
	for _, test := range tests {
		if got := maxProfit4(test.inputPrices, test.inputFee); got != test.want {
			t.Errorf("714_maxProfit4(%v, %v) = %v, expected %v", test.inputPrices, test.inputFee, got, test.want)
		}
	}
}
