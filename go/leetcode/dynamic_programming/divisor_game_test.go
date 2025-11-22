package dynamicprogramming

import "testing"

func TestDivisorGame(t *testing.T) {
	testCases := []testCase[int, bool]{
		{2, true},
		{3, false},
		{4, true},
	}

	t.Run("First Solution", func(t *testing.T) {
		runAllTestCases(t, testCases, divisorGame1)
	})

	t.Run("Second Solution", func(t *testing.T) {
		runAllTestCases(t, testCases, divisorGame2)
	})

	t.Run("Third Solution", func(t *testing.T) {
		runAllTestCases(t, testCases, divisorGame)
	})
}

func BenchmarkDivisorGame(b *testing.B) {
	input := 127
	b.Run("First Solution", func(b *testing.B) {
		for b.Loop() {
			divisorGame1(input)
		}
	})

	b.Run("Second Solution", func(b *testing.B) {
		for b.Loop() {
			divisorGame2(input)
		}
	})

	b.Run("Third Solution", func(b *testing.B) {
		for b.Loop() {
			divisorGame(input)
		}
	})
}
