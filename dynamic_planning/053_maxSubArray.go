package dynamic_planning

import "leetcode/utils"

func maxSubArray(nums []int) int {
	sum := 0
	maxSum := nums[0]
	for i := 0; i < len(nums); i++ {
		if sum > 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		maxSum = utils.MaxInt(sum, maxSum)
	}
	return maxSum
}
