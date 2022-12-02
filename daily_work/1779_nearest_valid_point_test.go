package daily_work

import (
	"math"
	"testing"
)

func nearestValidPoint(x int, y int, points [][]int) int {
	res := -1
	minDis := math.MaxInt32
	for i, point := range points {
		if point[0] != x && point[1] != y {
			continue
		}
		dis := abs(point[0]-x) + abs(point[1]-y)
		if dis < minDis {
			minDis = dis
			res = i
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func TestNearestValidPoint(t *testing.T) {
	type Input struct {
		x, y   int
		points [][]int
	}

	type TestCase struct {
		Input
		want int
	}
	tests := []TestCase{
		{Input{x: 3, y: 4, points: [][]int{{1, 2}, {3, 1}, {2, 4}, {2, 3}, {4, 4}}}, 2},
		{Input{x: 3, y: 4, points: [][]int{{3, 4}}}, 0},
		{Input{x: 3, y: 4, points: [][]int{{2, 3}}}, -1},
	}
	for _, test := range tests {
		if got := nearestValidPoint(test.x, test.y, test.points); got != test.want {
			t.Errorf("nearestValidPoint(%v) = %v, expected %v", test.Input, got, test.want)
		}
	}
}
