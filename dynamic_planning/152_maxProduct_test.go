package leetcode

import "testing"

func TestMaxProduct(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{2, 3, -2, 4}, 6},
		{[]int{-2, 0, -1}, 0},
		{[]int{5, 6, -1, -6, -2}, 180},
	}
	for _, test := range tests {
		if got := maxProduct(test.input); got != test.want {
			t.Errorf("152_maxProduct(%v) = %v, expected %v", test.input, got, test.want)
		}
	}
}
