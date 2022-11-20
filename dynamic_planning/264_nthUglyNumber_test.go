package dynamic_planning

import (
	"leetcode/utils"
	"testing"
)

func nthUglyNumber(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	p2, p3, p5 := 1, 1, 1
	for i := 2; i <= n; i++ {
		x2, x3, x5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
		dp[i] = utils.MinInt(x2, x3, x5)
		if dp[i] == x2 {
			p2++
		}
		if dp[i] == x3 {
			p3++
		}
		if dp[i] == x5 {
			p5++
		}
	}
	return dp[n]
}

func TestNthUglyNumber(t *testing.T) {
	var tests = []struct {
		input int
		want  int
	}{
		{10, 12},
		{1, 1},
	}
	for _, test := range tests {
		if got := nthUglyNumber(test.input); got != test.want {
			t.Errorf("264_nthUglyNumber(%v) = %v, expected %v", test.input, got, test.want)
		}
	}
}
