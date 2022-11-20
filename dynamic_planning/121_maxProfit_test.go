package dynamic_planning

import (
	"leetcode/utils"
	"testing"
)

func maxProfit(prices []int) int {
	ans := 0
	s := prices[0]
	si := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < s {
			s = prices[i]
			si = i
		}
		if i > si {
			ans = utils.MaxInt(ans, prices[i]-s)
		}
	}
	return ans
}

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{1, 2}, 1},
	}
	for _, test := range tests {
		if got := maxProfit(test.input); got != test.want {
			t.Errorf("121_maxProfit(%v) = %v, expected %v", test.input, got, test.want)
		}
	}
}
