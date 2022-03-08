package leetcode

import "testing"

func TestGetRow(t *testing.T) {
	tests := []struct {
		input int
		want  []int
	}{
		{3, []int{1, 3, 3, 1}},
		{0, []int{1}},
		{1, []int{1, 1}},
	}
	for _, test := range tests {
		got := getRow(test.input)
		flag := true
		if len(got) != len(test.want) {
			flag = false
		} else {
			for i, v := range test.want {
				if v != got[i] {
					flag = false
					break
				}
			}
		}
		if !flag {
			t.Errorf("119_getRow(%v) = %v, expected %v", test.input, got, test.want)
		}
	}
}
