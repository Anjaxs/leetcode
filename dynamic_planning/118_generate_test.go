package dynamic_planning

import "testing"

func generate(numRows int) [][]int {
	ans := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		ans[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				ans[i][j] = 1
			} else {
				ans[i][j] = ans[i-1][j-1] + ans[i-1][j]
			}
		}
	}
	return ans
}

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
