package main

import "testing"

func demo(a, b int) int {
	return a + b
}

func TestDemo(t *testing.T) {
	type Input struct {
		a, b int
	}
	type TestCase struct {
		Input
		want int
	}
	tests := []TestCase{
		{Input{1, 1}, 0},
		{Input{2, 1}, 0},
	}
	for _, test := range tests {
		if got := demo(test.a, test.b); got != test.want {
			t.Errorf("demo(%v) = %v, expected %v", test.Input, got, test.want)
		}
	}
}
