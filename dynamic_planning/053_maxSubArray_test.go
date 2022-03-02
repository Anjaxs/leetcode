package leetcode

import "testing"

func TestMaxSubArray(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{1}, 1},
		{[]int{5, 4, -1, 7, 8}, 23},
	}
	for _, test := range tests {
		if got := maxSubArray(test.input); got != test.want {
			t.Errorf("053_maxSubArray(%v) = %v, expected %v", test.input, got, test.want)
		}
	}
}
