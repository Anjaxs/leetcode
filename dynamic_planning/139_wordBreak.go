package leetcode

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
