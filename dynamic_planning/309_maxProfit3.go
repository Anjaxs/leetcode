package dynamic_planning

import "leetcode/utils"

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
