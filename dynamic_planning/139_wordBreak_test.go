package dynamic_planning

import "testing"

func wordBreak(s string, wordDict []string) bool {
	flag := make([]int, len(s)+1)
	dp(s, wordDict, 0, flag)
	return flag[len(s)] == 1
}

func dp(s string, wordDict []string, i int, flag []int) bool {
	if flag[i] == 1 {
		return true
	}
	if flag[i] == 2 {
		return false
	}
	if len(s) == 0 {
		flag[i] = 1
		return true
	}
	for _, v := range wordDict {
		if len(v) > len(s) {
			continue
		}
		if s[:len(v)] == v && dp(s[len(v):], wordDict, i+len(v), flag) {
			flag[i] = 1
			return true
		}
	}
	flag[i] = 2
	return false
}

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
