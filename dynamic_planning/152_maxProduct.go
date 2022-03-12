package dynamic_planning

import (
	"leetcode/utils"
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
