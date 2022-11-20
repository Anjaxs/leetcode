package dynamic_planning

import "testing"

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}

func TestChange(t *testing.T) {
	tests := []struct {
		amount int
		coins  []int
		want   int
	}{
		{5, []int{1, 2, 5}, 4},
		{3, []int{2}, 0},
	}
	for _, test := range tests {
		if got := change(test.amount, test.coins); got != test.want {
			t.Errorf("518_change(%v, %v) = %v, expected %v", test.coins, test.amount, got, test.want)
		}
	}
}
