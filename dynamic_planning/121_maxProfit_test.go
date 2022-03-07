package leetcode

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{1, 2}, 1},
	}
	for _, test := range tests {
		if got := maxProfit(test.input); got != test.want {
			t.Errorf("121_maxProfit(%v) = %v, expected %v", test.input, got, test.want)
		}
	}
}
