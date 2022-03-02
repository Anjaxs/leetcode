package leetcode

import "testing"

func TestGetMaxLen(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{1, -2, -3, 4}, 4},
		{[]int{0, 1, -2, -3, -4}, 3},
		{[]int{-1, -2, -3, 0, 1}, 2},
	}
	for _, test := range tests {
		if got := getMaxLen(test.input); got != test.want {
			t.Errorf("1567_getMaxLen(%v) = %v, expected %v", test.input, got, test.want)
		}
	}
}
