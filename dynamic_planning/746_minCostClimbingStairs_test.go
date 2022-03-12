package dynamic_planning

import "testing"

func TestMinCostClimbingStairs(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{10, 15, 20}, 15},
		{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
	}
	for _, test := range tests {
		if got := minCostClimbingStairs(test.input); got != test.want {
			t.Errorf("746_minCostClimbingStairs(%v) = %v, expected %v", test.input, got, test.want)
		}
	}
}
