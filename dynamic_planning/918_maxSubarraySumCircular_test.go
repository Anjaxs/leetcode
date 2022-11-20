package dynamic_planning

import (
	"leetcode/utils"
	"testing"
)

func maxSubarraySumCircular(nums []int) int {
	total, maxSum, minSum, currMax, currMin := nums[0], nums[0], nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		total += nums[i]
		currMax = utils.MaxInt(currMax+nums[i], nums[i])
		maxSum = utils.MaxInt(maxSum, currMax)
		currMin = utils.MinInt(currMin+nums[i], nums[i])
		minSum = utils.MinInt(minSum, currMin)
	}

	//等价于if maxSum < 0
	if total == minSum {
		return maxSum
	} else {
		return utils.MaxInt(maxSum, total-minSum)
	}
}

func TestMaxSubarraySumCircular(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{1}, 1},
		{[]int{5, 4, -1, 7, 8}, 24},
	}
	for _, test := range tests {
		if got := maxSubarraySumCircular(test.input); got != test.want {
			t.Errorf("918_maxSubarraySumCircular(%v) = %v, expected %v", test.input, got, test.want)
		}
	}
}
