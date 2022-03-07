package leetcode

import "leetcode/utils"

func maxProfit2(prices []int) int {
	n := len(prices)
	dp0, dp1 := 0, -prices[0]
	for i := 1; i < n; i++ {
		dp0, dp1 = utils.MaxInt(dp0, dp1+prices[i]), utils.MaxInt(dp1, dp0-prices[i])
	}
	return dp0
}
