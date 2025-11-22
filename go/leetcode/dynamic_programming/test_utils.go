package dynamicprogramming

import (
	"fmt"
	"slices"
	"testing"
)

func fmtMsgFunc(format string, args ...any) func() string {
	return func() string {
		return fmt.Sprintf(format, args...)
	}
}

func assert(t *testing.T, expr bool, msg string) {
	if expr != true {
		t.Error(msg)
	}
}

func assertFunc(t *testing.T, expr bool, msgFunc func() string) {
	if expr != true {
		t.Error(msgFunc())
	}
}

func assertEquals[A comparable](t *testing.T, expected, actual A) {
	assertFunc(t, expected == actual, fmtMsgFunc("Failed: %v != %v", expected, actual))
}

func assertArraysEqual[A comparable](t *testing.T, expected, actual []A) {
	assertFunc(t, slices.Equal(expected, actual), fmtMsgFunc("Failed:\n - Expected: %v\n - Actual: %v", expected, actual))
}

func assertTrue(t *testing.T, v bool) {
	assert(t, v == true, "Failed: Expected true, but got false")
}

func assertFalse(t *testing.T, v bool) {
	assert(t, v == false, "Failed: Expected false, but got true")
}

type testCase[I, O any] struct {
	input    I
	expected O
}

func runAllTestCasesFunc[I, O any](
	t *testing.T,
	testCases []testCase[I, O],
	fn func(I) O,
	assertion func(*testing.T, O, O)) {
	for i, tc := range testCases {
		testName := fmt.Sprintf("Test Case %d", i)
		t.Run(testName, func(t *testing.T) {
			assertion(t, tc.expected, fn(tc.input))
		})
	}
}

func runAllTestCases[I any, O comparable](t *testing.T, testCases []testCase[I, O], fn func(I) O) {
	for i, tc := range testCases {
		testName := fmt.Sprintf("Test Case %d", i)
		t.Run(testName, func(t *testing.T) {
			assertEquals(t, tc.expected, fn(tc.input))
		})
	}
}

func runAllTestsFunc[I any](t *testing.T, testCases []I, fn func(I) bool, msgFunc func(I) string) {
	for i, tc := range testCases {
		testName := fmt.Sprintf("Test Case %d", i)
		t.Run(testName, func(t *testing.T) {
			assertFunc(t, fn(tc), func() string {
				return msgFunc(tc)
			})
		})
	}
}
