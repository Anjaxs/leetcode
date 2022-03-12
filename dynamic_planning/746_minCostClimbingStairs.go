package dynamic_planning

import "leetcode/utils"

func minCostClimbingStairs(cost []int) int {
	// dp 代表到达 n 层需要的费用
	dp := make([]int, len(cost)+1)
	dp[0], dp[1] = 0, 0
	for i := 2; i <= len(cost); i++ {
		dp[i] = utils.MinInt(cost[i-1]+dp[i-1], cost[i-2]+dp[i-2])
	}
	return dp[len(cost)]
}
