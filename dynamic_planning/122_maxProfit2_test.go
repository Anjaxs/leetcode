package dynamic_planning

import (
	"leetcode/utils"
	"testing"
)

func maxProfit2(prices []int) int {
	n := len(prices)
	dp0, dp1 := 0, -prices[0]
	for i := 1; i < n; i++ {
		dp0, dp1 = utils.MaxInt(dp0, dp1+prices[i]), utils.MaxInt(dp1, dp0-prices[i])
	}
	return dp0
}

func TestMaxProfit2(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 7},
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{7, 6, 4, 3, 1}, 0},
	}
	for _, test := range tests {
		if got := maxProfit2(test.input); got != test.want {
			t.Errorf("122_maxProfit2(%v) = %v, expected %v", test.input, got, test.want)
		}
	}
}
