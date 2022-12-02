package daily_work

import (
	"leetcode/utils"
	"testing"
)

func minOperations(boxes string) []int {
	l, n := len(boxes), 0
	ans := make([]int, l)
	for i := 0; i < l; i++ {
		if boxes[i] == '1' {
			n++
			ans[0] += i
		}
	}
	m := 0
	for i := 1; i < l; i++ {
		if boxes[i-1] == '1' {
			m++
			n--
		}
		ans[i] = ans[i-1] - n + m
	}
	return ans
}

func TestMinOperations(t *testing.T) {
	type TestCase struct {
		boxes string
		want  []int
	}
	tests := []TestCase{
		{"001011", []int{11, 8, 5, 4, 3, 4}},
		{"110", []int{1, 1, 3}},
	}
	for _, test := range tests {
		if got := minOperations(test.boxes); !utils.IntSliceEqual(got, test.want) {
			t.Errorf("minOperations(%v) = %v, expected %v", test.boxes, got, test.want)
		}
	}
}
