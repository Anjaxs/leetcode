package dynamic_planning

import "leetcode/utils"

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
