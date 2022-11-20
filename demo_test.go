package main

import "testing"

func demo(a, b int) int {
	return 0
}

type Input struct {
	a, b int
}

type TestCase struct {
	input Input
	want  int
}

func TestDemo(t *testing.T) {
	tests := []TestCase{
		{Input{1, 1}, 0},
		{Input{2, 1}, 2},
	}
	for _, test := range tests {
		if got := demo(test.input.a, test.input.b); got != test.want {
			t.Errorf("demo(%v) = %v, expected %v", test.input, got, test.want)
		}
	}
}
