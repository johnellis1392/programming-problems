package dynamicprogramming

import (
	"fmt"
	"testing"
)

func TestIsSubsequence(t *testing.T) {
	testCases := []testCase[[]string, bool]{
		{[]string{"abc", "ahbgdc"}, true},
		{[]string{"axc", "ahbgdc"}, false},
		{[]string{"bb", "ahbgdc"}, false},
	}

	runAllTestsFunc(t, testCases, func(tc testCase[[]string, bool]) bool {
		return isSubsequence(tc.input[0], tc.input[1]) == tc.expected
	}, func(tc testCase[[]string, bool]) string {
		return fmt.Sprintf("Failed: Expected %v, but got %v", tc.expected, !tc.expected)
	})
}

func BenchmarkIsSubsequence(b *testing.B) {
	for b.Loop() {
		isSubsequence("abc", "ahbgdc")
	}
}
