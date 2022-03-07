package leetcode

import "testing"

func TestMaxProfit4(t *testing.T) {
	tests := []struct {
		inputPrices []int
		inputFee    int
		want        int
	}{
		{[]int{1, 3, 2, 8, 4, 9}, 2, 8},
		{[]int{1, 3, 7, 5, 10, 3}, 3, 6},
	}
	for _, test := range tests {
		if got := maxProfit4(test.inputPrices, test.inputFee); got != test.want {
			t.Errorf("714_maxProfit4(%v, %v) = %v, expected %v", test.inputPrices, test.inputFee, got, test.want)
		}
	}
}
