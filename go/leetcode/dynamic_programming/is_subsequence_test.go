package dynamicprogramming

import "testing"

func TestIsSubsequence(t *testing.T) {
	assertTrue := func(a bool) {
		if a != true {
			t.Errorf("Expected true, but got false")
		}
	}

	assertFalse := func(a bool) {
		if a != false {
			t.Errorf("Expected false, but got true")
		}
	}

	t.Run("Example 1", func(t *testing.T) {
		assertTrue(isSubsequence("abc", "ahbgdc"))
	})

	t.Run("Example 2", func(t *testing.T) {
		assertFalse(isSubsequence("axc", "ahbgdc"))
	})

	t.Run("Example 3", func(t *testing.T) {
		assertFalse(isSubsequence("bb", "ahbgdc"))
	})
}

func BenchmarkIsSubsequence(b *testing.B) {
	for b.Loop() {
		isSubsequence("abc", "ahbgdc")
	}
}
