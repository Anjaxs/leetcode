package dynamic_planning

func isSubsequence(s string, t string) bool {
	currIdx := 0
	for i := range t {
		if currIdx < len(s) && s[currIdx] == t[i] {
			currIdx++
		}
	}
	return currIdx == len(s)
}
