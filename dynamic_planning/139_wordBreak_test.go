package dynamic_planning

import "testing"

func TestWorkBreak(t *testing.T) {
	var tests = []struct {
		inputStr      string
		inputWordDict []string
		want          bool
	}{
		{"leetcode", []string{"leet", "code"}, true},
		{"applepenapple", []string{"apple", "pen"}, true},
		{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}, false},
	}
	for _, test := range tests {
		if got := wordBreak(test.inputStr, test.inputWordDict); got != test.want {
			t.Errorf("139_wordBreak(%v, %v) = %v, expected %v", test.inputStr, test.inputWordDict, got, test.want)
		}
	}
}
