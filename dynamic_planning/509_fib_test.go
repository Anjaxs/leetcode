package leetcode

import "testing"

func TestFib(t *testing.T) {
	var tests = []struct {
		input int
		want  int
	}{
		{2, 1},
		{3, 2},
		{4, 3},
	}
	for _, test := range tests {
		if got := fib(test.input); got != test.want {
			t.Errorf("509_fib(%v) = %v, expected %v", test.input, got, test.want)
		}
	}
}
