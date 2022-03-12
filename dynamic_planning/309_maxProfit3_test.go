package dynamic_planning

import "testing"

func TestMaxProfit3(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{1, 2, 3, 0, 2}, 3},
		{[]int{1}, 0},
	}
	for _, test := range tests {
		if got := maxProfit3(test.input); got != test.want {
			t.Errorf("309_maxProfit3(%v) = %v, expected %v", test.input, got, test.want)
		}
	}
}
