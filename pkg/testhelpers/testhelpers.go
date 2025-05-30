package testhelpers

import (
	"errors"
	"testing"
)

// AssertNotNilError fails the test if err is nil.
func AssertNotNilError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Errorf("expected error to be not nil, but got nil")
	}
}

// AssertNilError fails the test if err is not nil.
func AssertNilError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("expected error to be nil, but got not nill")
	}
}

// AssertEqualStrings fails the test if the input strings are not equal.
func AssertEqualStrings(t *testing.T, act, exp string) {
	t.Helper()
	if act != exp {
		t.Errorf("exp: %v, act: %v", exp, act)
	}
}

// AssertEqualInts fails the test if the input integers are not equal.
func AssertEqualInts(t *testing.T, act, exp int) {
	t.Helper()
	if act != exp {
		t.Errorf("exp: %v, act: %v", exp, act)
	}
}

// AssertEqualIntSlices fails the test if the input slices are not equal.
func AssertEqualIntSlices(t *testing.T, act, exp []int) {
	t.Helper()

	actLen, expLen := len(act), len(exp)
	if actLen != expLen {
		t.Errorf("exp arr len: %v, act arr len: %v", expLen, actLen)
	}

	for i := 0; i < expLen; i++ {
		if act[i] != exp[i] {
			t.Errorf("exp: %v, act: %v", exp, act)
		}
	}
}

// AssertCorrectError fails the test if error messages are not equal.
func AssertCorrectError(t *testing.T, act, exp error) {
	t.Helper()
	if !errors.Is(act, exp) {
		t.Errorf("exp: %v, act: %v", exp, act)
	}
}
