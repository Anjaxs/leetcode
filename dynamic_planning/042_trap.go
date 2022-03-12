package dynamic_planning

import "leetcode/utils"

func trap(height []int) (ans int) {
	n := len(height)
	if n == 0 {
		return
	}

	leftMax := make([]int, n)
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = utils.MaxInt(leftMax[i-1], height[i])
	}

	rightMax := make([]int, n)
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = utils.MaxInt(rightMax[i+1], height[i])
	}

	for i, h := range height {
		ans += utils.MinInt(leftMax[i], rightMax[i]) - h
	}
	return
}
