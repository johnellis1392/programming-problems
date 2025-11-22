package dynamicprogramming

import (
	"testing"
)

func TestClimbStairs(t *testing.T) {
	runAllTestCases(t, []testCase[int, int]{{2, 2}, {3, 3}}, ClimbStairs)
}
