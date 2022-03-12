package dynamic_planning

import "testing"

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
