package dynamicprogramming

import "testing"

func TestCountBits(t *testing.T) {
	testCases := []testCase[int, []int]{
		{2, []int{0, 1, 1}},
		{5, []int{0, 1, 1, 2, 1, 2}},
		{7, []int{0, 1, 1, 2, 1, 2, 2, 3}},
		{16, []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4, 1}},
	}
	for _, tc := range testCases {
		assertArraysEqual(t, tc.expected, countBits(tc.input))
	}
}

func BenchmarkCountBits(b *testing.B) {
	input := 127
	b.Run("Benchmark new solution", func(b *testing.B) {
		for b.Loop() {
			countBits(input)
		}
	})

	b.Run("Benchmark old solution", func(b *testing.B) {
		for b.Loop() {
			countBits_OLD(input)
		}
	})
}
