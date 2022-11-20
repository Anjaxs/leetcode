package dynamic_planning

import "testing"

func climbStairs(n int) int {
	a := []int{0, 1, 2}
	if n < 3 {
		return a[n]
	}
	for i := 3; i <= n; i++ {
		a[0], a[1], a[2] = a[1], a[2], a[1]+a[2]
	}
	return a[2]
}

func TestClimbStairs(t *testing.T) {
	var tests = []struct {
		input int
		want  int
	}{
		{2, 2},
		{3, 3},
	}
	for _, test := range tests {
		if got := climbStairs(test.input); got != test.want {
			t.Errorf("070_ClimbStairs(%v) = %v, expected %v", test.input, got, test.want)
		}
	}
}
