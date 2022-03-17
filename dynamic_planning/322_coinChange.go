package dynamic_planning

import "leetcode/utils"

func coinChange(coins []int, amount int) int {
	// dp 数组的定义：当目标金额为 i 时，至少需要 dp[i] 枚硬币凑出
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}

	dp[0] = 0
	for i := range dp {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = utils.MinInt(dp[i], 1+dp[i-coin])
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
