package dynamic_planning

import (
	"leetcode/utils"
	"testing"
)

func getMaxLen(nums []int) (ans int) {
	pos, neg := 0, 0
	for _, num := range nums {
		if num > 0 {
			pos++
			if neg > 0 {
				neg++
			}
		} else if num == 0 {
			pos, neg = 0, 0
		} else {
			if neg > 0 {
				pos, neg = neg+1, pos+1
			} else {
				pos, neg = 0, pos+1
			}
		}
		ans = utils.MaxInt(ans, pos)
	}
	return
}

func TestGetMaxLen(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{1, -2, -3, 4}, 4},
		{[]int{0, 1, -2, -3, -4}, 3},
		{[]int{-1, -2, -3, 0, 1}, 2},
	}
	for _, test := range tests {
		if got := getMaxLen(test.input); got != test.want {
			t.Errorf("1567_getMaxLen(%v) = %v, expected %v", test.input, got, test.want)
		}
	}
}
