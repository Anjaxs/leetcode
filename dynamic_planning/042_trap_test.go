package dynamic_planning

import (
	"leetcode/utils"
	"testing"
)

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

func TestTrap(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{[]int{4, 2, 0, 3, 2, 5}, 9},
	}
	for _, test := range tests {
		if got := trap(test.input); got != test.want {
			t.Errorf("042_trap(%v) = %v, expected %v", test.input, got, test.want)
		}
	}
}
