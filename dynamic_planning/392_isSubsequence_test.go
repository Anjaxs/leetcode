package dynamic_planning

import "testing"

func isSubsequence(s string, t string) bool {
	currIdx := 0
	for i := range t {
		if currIdx < len(s) && s[currIdx] == t[i] {
			currIdx++
		}
	}
	return currIdx == len(s)
}

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want bool
	}{
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", false},
	}
	for _, test := range tests {
		if got := isSubsequence(test.s, test.t); got != test.want {
			t.Errorf("392_isSubsequence(%v, %v) = %v, expected %v", test.s, test.t, got, test.want)
		}
	}
}
