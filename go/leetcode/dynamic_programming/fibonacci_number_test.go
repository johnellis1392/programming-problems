package dynamicprogramming

import "testing"

func TestFibonacciNumber(t *testing.T) {
	assertEquals := func(expected, actual int) {
		if expected != actual {
			t.Errorf("Failed: %v != %v", actual, expected)
		}
	}

	t.Run("Example 1", func(t *testing.T) {
		assertEquals(1, fib(2))
	})

	t.Run("Example 2", func(t *testing.T) {
		assertEquals(2, fib(3))
	})

	t.Run("Example 3", func(t *testing.T) {
		assertEquals(3, fib(4))
	})
}

func BenchmarkFib(b *testing.B) {
	for b.Loop() {
		fib(10)
	}
}
