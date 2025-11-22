package dynamicprogramming

import "testing"

func TestFibonacciNumber(t *testing.T) {
	testCases := []testCase[int, int]{
		{2, 1},
		{3, 2},
		{4, 3},
	}
	runAllTestCases(t, testCases, fib)
}

func BenchmarkFib(b *testing.B) {
	for b.Loop() {
		fib(10)
	}
}
