package dynamic_planning

import "testing"

func tribonacci(n int) int {
	a := []int{0, 1, 1}
	if n <= 2 {
		return a[n]
	}
	for i := 2; i < n; i++ {
		a[0], a[1], a[2] = a[1], a[2], a[0]+a[1]+a[2]
	}
	return a[2]
}

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
