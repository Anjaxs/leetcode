package dynamic_planning

import "testing"

func TestMaxProfit2(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 7},
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{7, 6, 4, 3, 1}, 0},
	}
	for _, test := range tests {
		if got := maxProfit2(test.input); got != test.want {
			t.Errorf("122_maxProfit2(%v) = %v, expected %v", test.input, got, test.want)
		}
	}
}
