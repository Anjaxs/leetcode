package dynamic_planning

import "testing"

func TestTrap(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{[]int{4, 2, 0, 3, 2, 5}, 9},
	}
	for _, test := range tests {
		if got := trap(test.input); got != test.want {
			t.Errorf("042_trap(%v) = %v, expected %v", test.input, got, test.want)
		}
	}
}
