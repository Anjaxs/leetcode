package daily_work

import "testing"

func champagneTower(poured int, query_row int, query_glass int) float64 {
	glassTower := make([][]float64, query_row+1)
	for i := 0; i <= query_row; i++ {
		glassTower[i] = make([]float64, i+1)
	}
	glassTower[0][0] = float64(poured)
	for i := 1; i <= query_row; i++ {
		for j := 0; j <= i; j++ {
			if j > 0 && glassTower[i-1][j-1] > 1 {
				glassTower[i][j] += (glassTower[i-1][j-1] - 1) / 2
			}
			if j < i && glassTower[i-1][j] > 1 {
				glassTower[i][j] += (glassTower[i-1][j] - 1) / 2
			}
		}
	}
	if glassTower[query_row][query_glass] > 1 {
		return 1
	}
	return glassTower[query_row][query_glass]
}

type Input struct {
	poured, queryRow, queryGlass int
}

type TestCase struct {
	input Input
	want  float64
}

func TestChampagneTower(t *testing.T) {
	tests := []TestCase{
		{Input{1, 1, 1}, 0},
		{Input{2, 1, 1}, 0.5},
		{Input{100000009, 33, 17}, 1},
	}
	for _, test := range tests {
		if got := champagneTower(test.input.poured, test.input.queryRow, test.input.queryGlass); got != test.want {
			t.Errorf("champagneTower(%v) = %v, expected %v", test.input, got, test.want)
		}
	}
}
