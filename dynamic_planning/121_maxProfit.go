package dynamic_planning

import "leetcode/utils"

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
