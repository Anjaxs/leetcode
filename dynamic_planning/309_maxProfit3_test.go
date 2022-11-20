package dynamic_planning

import (
	"leetcode/utils"
	"testing"
)

func maxProfit3(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)
	f0, f1, f2 := -prices[0], 0, 0
	for i := 1; i < n; i++ {
		newf0 := utils.MaxInt(f0, f2-prices[i])
		newf1 := f0 + prices[i]
		newf2 := utils.MaxInt(f1, f2)
		f0, f1, f2 = newf0, newf1, newf2
	}
	return utils.MaxInt(f1, f2)
}

func TestMaxProfit3(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{1, 2, 3, 0, 2}, 3},
		{[]int{1}, 0},
	}
	for _, test := range tests {
		if got := maxProfit3(test.input); got != test.want {
			t.Errorf("309_maxProfit3(%v) = %v, expected %v", test.input, got, test.want)
		}
	}
}
