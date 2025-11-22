package dynamicprogramming

import "testing"

func TestMaxProfit(t *testing.T) {
	testCases := []testCase[[]int, int]{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
	}
	runAllTestCases(t, testCases, maxProfit)
}
