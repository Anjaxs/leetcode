package leetcode

import "testing"

func TestMaxScoreSightseeingPair(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{8, 1, 5, 2, 6}, 11},
		{[]int{1, 2}, 2},
		{[]int{5, 4, -1, 7, 8}, 14},
	}
	for _, test := range tests {
		if got := maxScoreSightseeingPair(test.input); got != test.want {
			t.Errorf("1014_maxScoreSightseeingPair(%v) = %v, expected %v", test.input, got, test.want)
		}
	}
}
