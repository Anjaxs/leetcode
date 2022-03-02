package leetcode

import "testing"

func TestTribonacci(t *testing.T) {
	var tests = []struct {
		input int
		want  int
	}{
		{4, 4},
		{25, 1389537},
	}
	for _, test := range tests {
		if got := tribonacci(test.input); got != test.want {
			t.Errorf("1137_tribonacci(%v) = %v, expected %v", test.input, got, test.want)
		}
	}
}
