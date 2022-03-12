package dynamic_planning

import "testing"

func TestGenerate(t *testing.T) {
	tests := []struct {
		input int
		want  [][]int
	}{
		{5, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
		{1, [][]int{{1}}},
	}
	for _, test := range tests {
		got := generate(test.input)
		flag := true
		if len(got) != len(test.want) {
			flag = false
		} else {
			for i, v := range test.want {
				if len(v) != len(got[i]) {
					flag = false
					break
				}
				for j, v2 := range v {
					if v2 != got[i][j] {
						flag = false
						break
					}
				}
			}
		}
		if !flag {
			t.Errorf("118_generate(%v) = %v, expected %v", test.input, got, test.want)
		}
	}
}
