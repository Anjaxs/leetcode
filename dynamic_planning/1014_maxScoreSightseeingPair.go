package dynamic_planning

import "leetcode/utils"

func maxScoreSightseeingPair(values []int) int {
	ans, mx := 0, values[0]+0
	for j := 1; j < len(values); j++ {
		ans = utils.MaxInt(ans, mx+values[j]-j)
		// 边遍历边维护
		mx = utils.MaxInt(mx, values[j]+j)
	}
	return ans
}
