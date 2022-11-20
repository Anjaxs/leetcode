package dynamic_planning

import "testing"

func getRow(rowIndex int) []int {
	prevRow := make([]int, rowIndex+1)
	row := make([]int, rowIndex+1)
	for i := 0; i < rowIndex+1; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				row[j] = 1
			} else {
				row[j] = prevRow[j-1] + prevRow[j]
			}
		}
		copy(prevRow, row)
	}
	return row
}

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
