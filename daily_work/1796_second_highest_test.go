package daily_work

import (
	"testing"
	"unicode"
)

func secondHighest(s string) int {
	a, b := -1, -1
	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			c := int(s[i] - '0')
			if c > a {
				a, b = c, a
			} else if c == a {
				continue
			} else if c > b {
				b = c
			}
		}
	}
	return b
}

func TestSecondHighest(t *testing.T) {
	type TestCase struct {
		input string
		want  int
	}
	tests := []TestCase{
		{"dfa12321afd", 2},
		{"abc1111", -1},
		{"ck077", 0},
		{"sjhtz8344", 4},
	}
	for _, test := range tests {
		if got := secondHighest(test.input); got != test.want {
			t.Errorf("secondHighest(%v) = %v, expected %v", test.input, got, test.want)
		}
	}
}
