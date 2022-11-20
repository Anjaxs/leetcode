package dynamic_planning

import (
	"leetcode/utils"
	"testing"
)

func minCostClimbingStairs(cost []int) int {
	// dp 代表到达 n 层需要的费用
	dp := make([]int, len(cost)+1)
	dp[0], dp[1] = 0, 0
	for i := 2; i <= len(cost); i++ {
		dp[i] = utils.MinInt(cost[i-1]+dp[i-1], cost[i-2]+dp[i-2])
	}
	return dp[len(cost)]
}

func TestMinCostClimbingStairs(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{10, 15, 20}, 15},
		{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
	}
	for _, test := range tests {
		if got := minCostClimbingStairs(test.input); got != test.want {
			t.Errorf("746_minCostClimbingStairs(%v) = %v, expected %v", test.input, got, test.want)
		}
	}
}
