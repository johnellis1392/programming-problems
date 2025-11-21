package dynamicprogramming

import (
	"testing"
)

func TestClimbStairs(t *testing.T) {
	assertEquals := func(expected, actual int) {
		if expected != actual {
			t.Errorf("Failed: %v != %v", actual, expected)
		}
	}

	t.Run("Example 1", func(t *testing.T) {
		assertEquals(2, ClimbStairs(2))
	})

	t.Run("Example 2", func(t *testing.T) {
		assertEquals(3, ClimbStairs(3))
	})
}
