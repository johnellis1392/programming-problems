package dynamicprogramming

import "testing"

func TestTribonacci(t *testing.T) {
	testCases := []testCase[int, int]{
		{4, 4},
		{25, 1389537},
	}
	runAllTestCases(t, testCases, tribonacci)
}

func BenchmarkTribonacci(b *testing.B) {
	b.Run("Benchmark", func(b *testing.B) {
		for b.Loop() {
			tribonacci(25)
		}
	})
}
