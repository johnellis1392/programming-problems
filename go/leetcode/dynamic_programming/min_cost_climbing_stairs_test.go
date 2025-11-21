package dynamicprogramming

import "testing"

func TestMinCostClimbingStairs(t *testing.T) {
	assertEquals := func(expected, actual int) {
		if expected != actual {
			t.Errorf("Failed: %v != %v", actual, expected)
		}
	}

	t.Run("Example 1", func(t *testing.T) {
		input := []int{10, 15, 20}
		expected := 15
		assertEquals(expected, minCostClimbingStairs(input))
	})

	t.Run("Example 2", func(t *testing.T) {
		input := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
		expected := 6
		assertEquals(expected, minCostClimbingStairs(input))
	})
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
