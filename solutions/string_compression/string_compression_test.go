package main

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []string{
		"abcabcbb",
		"aaaaaaaa",
		"abcde",
		"",
		"pwwkew",
	}
	results := []string{
		"abcabcbb",
		"a8",
		"abcde",
		"",
		"pwwkew",
	}
	caseNum := 5
	for i := 0; i < caseNum; i++ {
		if ret := compressString(tests[i]); ret != results[i] {
			t.Fatalf("case %d failed\nactual: %s, expect: %s\n", i, ret, results[i])
		}
	}
}
