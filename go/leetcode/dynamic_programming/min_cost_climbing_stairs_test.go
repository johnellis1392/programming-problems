package dynamicprogramming

import "testing"

func TestMinCostClimbingStairs(t *testing.T) {
	testCases := []testCase[[]int, int]{
		{[]int{10, 15, 20}, 15},
		{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
	}
	runAllTestCases(t, testCases, minCostClimbingStairs)
}

func BenchmarkMinCostClimbingStairs(b *testing.B) {
	input := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	b.Run("Benchmark Old implementation", func(b *testing.B) {
		for b.Loop() {
			minCostClimbingStairs_OLD(input)
		}
	})

	b.Run("Benchmark New implementation", func(b *testing.B) {
		for b.Loop() {
			minCostClimbingStairs(input)
		}
	})
}
