package dynamicprogramming

import "testing"

func TestPascalsTriangle2(t *testing.T) {
	testCases := []testCase[int, []int]{
		{3, []int{1, 3, 3, 1}},
		{0, []int{1}},
		{1, []int{1, 1}},
	}
	runAllTestCasesFunc(t, testCases, getRow, assertArraysEqual)
}
