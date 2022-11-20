package dynamic_planning

import (
	"leetcode/utils"
	"testing"
)

func maxProduct(nums []int) int {
	maxF, minF, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxF, minF
		maxF = utils.MaxInt(mx*nums[i], utils.MaxInt(nums[i], mn*nums[i]))
		minF = utils.MinInt(mn*nums[i], utils.MinInt(nums[i], mx*nums[i]))
		ans = utils.MaxInt(maxF, ans)
	}
	return ans
}

func TestMaxProduct(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{2, 3, -2, 4}, 6},
		{[]int{-2, 0, -1}, 0},
		{[]int{5, 6, -1, -6, -2}, 180},
	}
	for _, test := range tests {
		if got := maxProduct(test.input); got != test.want {
			t.Errorf("152_maxProduct(%v) = %v, expected %v", test.input, got, test.want)
		}
	}
}
