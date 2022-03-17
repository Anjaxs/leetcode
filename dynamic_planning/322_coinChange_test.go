package dynamic_planning

import "testing"

func TestCoinChange(t *testing.T) {
	tests := []struct {
		coins  []int
		amount int
		want   int
	}{
		{[]int{1, 2, 5}, 11, 3},
		{[]int{2}, 3, -1},
	}
	for _, test := range tests {
		if got := coinChange(test.coins, test.amount); got != test.want {
			t.Errorf("322_coinChange(%v, %v) = %v, expected %v", test.coins, test.amount, got, test.want)
		}
	}
}
