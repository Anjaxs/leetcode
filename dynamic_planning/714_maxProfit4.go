package dynamic_planning

import "leetcode/utils"

func maxProfit4(prices []int, fee int) int {
	n := len(prices)
	dp0, dp1 := 0, -prices[0]
	for i := 1; i < n; i++ {
		dp0, dp1 = utils.MaxInt(dp0, dp1+prices[i]-fee), utils.MaxInt(dp1, dp0-prices[i])
	}
	return utils.MaxInt(dp0, dp1)
}
